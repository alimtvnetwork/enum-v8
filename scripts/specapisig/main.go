// Command specapisig — Go AST signature indexer for S-106 v2.
//
// Walks one or more Go module roots, parses every non-test .go file with
// go/parser, and emits a JSON signature index of every exported top-level
// function keyed by `package.Symbol`. Methods include their receiver type.
//
// Designed for the PowerShell driver scripts/spec-api-sig-check.psm1
// which consumes the JSON to validate call-site arity and parameter
// kinds against spec markdown snippets.
//
// Exit codes:
//
//	0 — index written successfully.
//	1 — at least one root failed to parse (logged to stderr).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SigItem struct {
	Package  string  `json:"package"`
	Symbol   string  `json:"symbol"`
	Kind     string  `json:"kind"` // "func" | "method"
	Receiver string  `json:"receiver,omitempty"`
	Params   []Param `json:"params"`
	Results  []Param `json:"results"`
	Variadic bool    `json:"variadic"`
	File     string  `json:"file"`
	Line     int     `json:"line"`
}

type Index struct {
	Version   string    `json:"version"`
	Generated string    `json:"generated"`
	Roots     []string  `json:"roots"`
	Packages  int       `json:"packages"`
	Functions int       `json:"functions"`
	Items     []SigItem `json:"items"`
}

const indexVersion = "1.0.0"

func main() {
	rootsFlag := flag.String("roots", "/tmp/core-v9-upstream,.",
		"comma-separated directory roots to index")
	outFlag := flag.String("out", "/tmp/core-v9-sigindex.json",
		"output JSON path")
	skipFlag := flag.String("skip",
		`^(\.git|node_modules|vendor|cross-repo|tests|scripts|spec|src|public|data|cmd|assets|configs|internal|dist|build|coverage|\.lovable|\.github|\.ci-baselines|\.ci-cache|\.release)$`,
		"regex of directory basenames to skip")
	flag.Parse()

	skipRe, err := regexp.Compile(*skipFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid -skip regex: %v\n", err)
		os.Exit(2)
	}

	roots := splitCSV(*rootsFlag)
	idx := Index{
		Version:   indexVersion,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Roots:     roots,
		Items:     []SigItem{},
	}

	pkgSeen := map[string]bool{}
	hadError := false
	fset := token.NewFileSet()

	for _, root := range roots {
		absRoot, _ := filepath.Abs(root)
		err := filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil // tolerate transient errors
			}
			if d.IsDir() {
				if skipRe.MatchString(d.Name()) {
					return filepath.SkipDir
				}
				return nil
			}
			if !strings.HasSuffix(d.Name(), ".go") {
				return nil
			}
			if strings.HasSuffix(d.Name(), "_test.go") {
				return nil
			}
			file, perr := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
			if perr != nil {
				// Don't fail the run on single-file parse errors — log and continue.
				fmt.Fprintf(os.Stderr, "  parse warn: %s: %v\n", path, perr)
				return nil
			}
			pkg := file.Name.Name
			pkgSeen[pkg] = true

			rel, _ := filepath.Rel(absRoot, path)
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}
				if !fn.Name.IsExported() {
					continue
				}
				item := SigItem{
					Package: pkg,
					Symbol:  fn.Name.Name,
					Kind:    "func",
					Params:  flattenFields(fn.Type.Params),
					Results: flattenFields(fn.Type.Results),
					File:    rel,
					Line:    fset.Position(fn.Pos()).Line,
				}
				if fn.Recv != nil && len(fn.Recv.List) > 0 {
					item.Kind = "method"
					item.Receiver = exprString(fn.Recv.List[0].Type)
				}
				if fn.Type.Params != nil {
					for _, f := range fn.Type.Params.List {
						if _, isEll := f.Type.(*ast.Ellipsis); isEll {
							item.Variadic = true
						}
					}
				}
				idx.Items = append(idx.Items, item)
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "walk error in %s: %v\n", absRoot, err)
			hadError = true
		}
	}

	idx.Packages = len(pkgSeen)
	idx.Functions = len(idx.Items)

	out, err := os.Create(*outFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(idx); err != nil {
		fmt.Fprintf(os.Stderr, "encode: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ specapisig v%s — %d packages, %d exported funcs/methods → %s\n",
		indexVersion, idx.Packages, idx.Functions, *outFlag)

	if hadError {
		os.Exit(1)
	}
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// flattenFields turns an *ast.FieldList (where one Names slice may bind
// multiple identifiers to the same Type, e.g. `a, b int`) into a flat
// list of {Name, Type} entries.
func flattenFields(fl *ast.FieldList) []Param {
	out := []Param{}
	if fl == nil {
		return out
	}
	for _, f := range fl.List {
		typ := exprString(f.Type)
		if len(f.Names) == 0 {
			// Unnamed result, e.g. `func Foo() error`.
			out = append(out, Param{Name: "", Type: typ})
			continue
		}
		for _, n := range f.Names {
			out = append(out, Param{Name: n.Name, Type: typ})
		}
	}
	return out
}

// exprString renders a Go type expression as a stable string suitable for
// equality comparison. Handles pointers, slices, maps, channels, function
// types, ellipsis (variadic), selectors (`pkg.T`), and identifiers.
func exprString(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprString(t.X)
	case *ast.SelectorExpr:
		return exprString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + exprString(t.Elt)
		}
		return "[" + exprString(t.Len) + "]" + exprString(t.Elt)
	case *ast.MapType:
		return "map[" + exprString(t.Key) + "]" + exprString(t.Value)
	case *ast.ChanType:
		return "chan " + exprString(t.Value)
	case *ast.Ellipsis:
		return "..." + exprString(t.Elt)
	case *ast.InterfaceType:
		if t.Methods == nil || len(t.Methods.List) == 0 {
			return "interface{}"
		}
		return "interface{...}"
	case *ast.FuncType:
		return "func(" + joinParams(t.Params) + ")" + resultStr(t.Results)
	case *ast.StructType:
		return "struct{...}"
	case *ast.BasicLit:
		return t.Value
	case *ast.ParenExpr:
		return "(" + exprString(t.X) + ")"
	default:
		return fmt.Sprintf("%T", e)
	}
}

func joinParams(fl *ast.FieldList) string {
	parts := []string{}
	if fl == nil {
		return ""
	}
	for _, f := range fl.List {
		parts = append(parts, exprString(f.Type))
	}
	return strings.Join(parts, ",")
}

func resultStr(fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return ""
	}
	if len(fl.List) == 1 && len(fl.List[0].Names) == 0 {
		return " " + exprString(fl.List[0].Type)
	}
	return " (" + joinParams(fl) + ")"
}