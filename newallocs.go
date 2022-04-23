package consistent

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	newAllocsLiteral = "literal"
	newAllocsNew     = "new"
)

var newAllocsFlagAllowedValues = []string{flagIgnore, newAllocsLiteral, newAllocsNew}

func checkNewAllocLit(pass *analysis.Pass, expr *ast.UnaryExpr, mode string) {
	if mode != newAllocsNew {
		return
	}

	if expr.Op != token.AND {
		return
	}

	comp := astcast.ToCompositeLit(expr.X)

	if !identOrSelector(comp.Type) {
		return
	}

	if len(comp.Elts) != 0 {
		return
	}

	pass.Reportf(expr.Pos(), "call new instead of using zero-value literal")
}

func checkNewAllocNew(pass *analysis.Pass, call *ast.CallExpr, mode string) {
	if mode != newAllocsLiteral {
		return
	}

	if astcast.ToIdent(call.Fun).Name != "new" {
		return
	}

	if len(call.Args) != 1 {
		return
	}

	pass.Reportf(call.Pos(), "use zero-value literal instead of calling new")
}

func identOrSelector(e ast.Expr) bool {
	switch e.(type) {
	case *ast.Ident, *ast.SelectorExpr:
		return true
	default:
		return false
	}
}
