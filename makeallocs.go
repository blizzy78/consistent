package consistent

import (
	"go/ast"
	"go/types"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	makeAllocsLiteral = "literal"
	makeAllocsMake    = "make"
)

var makeAllocsFlagAllowedValues = []string{flagIgnore, makeAllocsLiteral, makeAllocsMake}

func checkMakeAllocLit(pass *analysis.Pass, expr *ast.CompositeLit, mode string) {
	if mode != makeAllocsMake {
		return
	}

	if len(expr.Elts) != 0 {
		return
	}

	switch pass.TypesInfo.TypeOf(expr.Type).Underlying().(type) {
	case *types.Slice:
		pass.Reportf(expr.Pos(), "call make instead of using slice literal")
	case *types.Map:
		pass.Reportf(expr.Pos(), "call make instead of using map literal")
	}
}

func checkMakeAllocMake(pass *analysis.Pass, call *ast.CallExpr, mode string) { //nolint:gocognit,cyclop // lots of constraints to check
	if mode != makeAllocsLiteral {
		return
	}

	if astcast.ToIdent(call.Fun).Name != "make" {
		return
	}

	if len(call.Args) == 0 {
		return
	}

	typ := pass.TypesInfo.TypeOf(call.Args[0]).Underlying()
	_, arrayTyp := typ.(*types.Slice)
	_, mapTyp := typ.(*types.Map)

	switch {
	case arrayTyp && (len(call.Args) < 2 || len(call.Args) > 3):
		return
	case mapTyp && len(call.Args) > 2:
		return
	case !arrayTyp && !mapTyp:
		return
	}

	if len(call.Args) >= 2 {
		if l, ok := litInt(call.Args[1]); !ok || l != 0 {
			return
		}
	}

	if arrayTyp && len(call.Args) == 3 {
		if c, ok := litInt(call.Args[2]); !ok || c != 0 {
			return
		}
	}

	if mapTyp {
		pass.Reportf(call.Pos(), "use map literal instead of calling make")
		return
	}

	pass.Reportf(call.Pos(), "use slice literal instead of calling make")
}
