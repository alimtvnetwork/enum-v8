package creationtests

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_BasicStringDefectDetector enforces RCA Pattern 3 across the
// repository: any package declaring `type Variant string` (the BasicString
// pattern) must override the canonical method set, because the upstream
// `core-v9 v1.5.8` `BasicString` implementation has known defects
// (PI-005..007 cluster) where these accessors return zero values for
// spread-constructed enums.
//
// Required overrides (declared on the package's Variant receiver):
//   - MinValueString()
//   - MarshalJSON()
//   - UnmarshalJSON()
//   - IsAnyNamesOf()
//   - RangesDynamicMap()
//
// Implementation: AST-walks every Go file in the repo (skipping vendored,
// cross-repo, tests, and node_modules paths). For each `type Variant string`
// declaration, scans the same package directory for method declarations
// whose receiver type is `Variant` (or `*Variant`).
//
// Failure points the author at the exact package and missing method names.
var basicStringDefectDetectorRequired = []string{
	"MinValueString",
	"MarshalJSON",
	"UnmarshalJSON",
	"IsAnyNamesOf",
	"RangesDynamicMap",
}

var basicStringDefectDetectorSkipDir = map[string]bool{
	"node_modules": true,
	"cross-repo":   true,
	"src":          true,
	".git":         true,
	".lovable":     true,
	"data":         true,
	"spec":         true,
	"tests":        true,
	"scripts":      true,
}

func Test_AllEnums_BasicStringDefectDetector(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Skipf("cannot locate repo root: %v", err)
		return
	}

	stringVariantPkgs := map[string]string{} // pkgDir -> import-relative name
	err = filepath.Walk(root, func(path string, info os.FileInfo, werr error) error {
		if werr != nil {
			return nil
		}
		if info.IsDir() {
			if basicStringDefectDetectorSkipDir[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		af, perr := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if perr != nil {
			return nil
		}
		for _, decl := range af.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || ts.Name.Name != "Variant" {
					continue
				}
				ident, ok := ts.Type.(*ast.Ident)
				if !ok || ident.Name != "string" {
					continue
				}
				stringVariantPkgs[filepath.Dir(path)] = af.Name.Name
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk failed: %v", err)
	}

	for pkgDir, pkgName := range stringVariantPkgs {
		pkgDir, pkgName := pkgDir, pkgName
		Convey(pkgName+" — string-backed Variant must override BasicString methods (RCA Pattern 3)", t, func() {
			found := map[string]bool{}
			entries, _ := os.ReadDir(pkgDir)
			for _, e := range entries {
				if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
					continue
				}
				fset := token.NewFileSet()
				af, perr := parser.ParseFile(fset, filepath.Join(pkgDir, e.Name()), nil, parser.SkipObjectResolution)
				if perr != nil {
					continue
				}
				for _, decl := range af.Decls {
					fd, ok := decl.(*ast.FuncDecl)
					if !ok || fd.Recv == nil || len(fd.Recv.List) == 0 {
						continue
					}
					recvType := fd.Recv.List[0].Type
					if star, ok := recvType.(*ast.StarExpr); ok {
						recvType = star.X
					}
					ident, ok := recvType.(*ast.Ident)
					if !ok || ident.Name != "Variant" {
						continue
					}
					found[fd.Name.Name] = true
				}
			}

			var missing []string
			for _, req := range basicStringDefectDetectorRequired {
				if !found[req] {
					missing = append(missing, req)
				}
			}
			So(missing, ShouldBeEmpty)
		})
	}
}

func findRepoRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := cwd
	for i := 0; i < 8; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", os.ErrNotExist
}
