package creationtests

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_LevelComparisonReceiverOrder enforces RCA Pattern 5 across
// every non-test source file: any method named `IsAboveOrEqual`,
// `IsAbove`, `IsLowerOrEqual`, or `IsLower` whose body is a single
// `return X OP Y` over `ValueByte()` calls must place the receiver on
// the LEFT of the operator (i.e. `it.ValueByte() OP arg.ValueByte()`).
//
// The historical defect: many packages copy-pasted `level.ValueByte() >=
// it.ValueByte()` which inverts the semantics — `Error.IsAboveOrEqual(
// Notice)` returns false instead of true. Fixed wholesale in v0.87.0.
//
// This guard prevents regression. AST-walks the repo, finds every method
// matching the names above, and asserts the LHS of the comparison is a
// `ValueByte()` call on the receiver identifier.
var levelCompareTargetMethods = map[string]bool{
	"IsAboveOrEqual": true,
	"IsAbove":        true,
	"IsLowerOrEqual": true,
	"IsLower":        true,
}

var levelCompareSkipDir = map[string]bool{
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

func Test_AllEnums_LevelComparisonReceiverOrder(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Skipf("cannot locate repo root: %v", err)
		return
	}

	type finding struct {
		file, method, recvName, body string
	}
	var inversions []finding

	_ = filepath.Walk(root, func(path string, info os.FileInfo, werr error) error {
		if werr != nil {
			return nil
		}
		if info.IsDir() {
			if levelCompareSkipDir[info.Name()] {
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
			fd, ok := decl.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || len(fd.Recv.List) == 0 {
				continue
			}
			if !levelCompareTargetMethods[fd.Name.Name] {
				continue
			}
			if len(fd.Recv.List[0].Names) == 0 {
				continue
			}
			recvName := fd.Recv.List[0].Names[0].Name
			if fd.Body == nil || len(fd.Body.List) != 1 {
				continue
			}
			ret, ok := fd.Body.List[0].(*ast.ReturnStmt)
			if !ok || len(ret.Results) != 1 {
				continue
			}
			bin, ok := ret.Results[0].(*ast.BinaryExpr)
			if !ok {
				continue
			}
			lhsRecv := receiverOfValueByteCall(bin.X)
			if lhsRecv == "" {
				continue // not a recognized form; ignore
			}
			if lhsRecv != recvName {
				var sb strings.Builder
				_ = printer.Fprint(&sb, fset, ret)
				inversions = append(inversions, finding{
					file:     path,
					method:   fd.Name.Name,
					recvName: recvName,
					body:     sb.String(),
				})
			}
		}
		return nil
	})

	Convey("Level-comparison methods must place receiver on LHS (RCA Pattern 5)", t, func() {
		if len(inversions) > 0 {
			for _, f := range inversions {
				t.Errorf("%s: %s has inverted receiver order — %s (expected %s.ValueByte() OP arg.ValueByte())", f.file, f.method, f.body, f.recvName)
			}
		}
		So(inversions, ShouldBeEmpty)
	})
}

// receiverOfValueByteCall returns the identifier name of `X` if expr is
// `X.ValueByte()` (a zero-arg method call), else "".
func receiverOfValueByteCall(expr ast.Expr) string {
	call, ok := expr.(*ast.CallExpr)
	if !ok || len(call.Args) != 0 {
		return ""
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "ValueByte" {
		return ""
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return ""
	}
	return id.Name
}
