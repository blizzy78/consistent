package consistent

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	andNOTsANDNOT  = "andNot"
	andNOTsANDComp = "andComp"
)

var andNOTsFlagAllowedValues = []string{flagIgnore, andNOTsANDNOT, andNOTsANDComp}

func checkAndNotExpr(pass *analysis.Pass, expr *ast.BinaryExpr, mode string) {
	switch mode {
	case flagIgnore:
		return

	case andNOTsANDNOT:
		if expr.Op != token.AND {
			return
		}

		if astcast.ToUnaryExpr(expr.Y).Op != token.XOR {
			return
		}

		pass.Reportf(expr.Pos(), "use AND-NOT operator instead of AND operator with complement expression")

	case andNOTsANDComp:
		if expr.Op != token.AND_NOT {
			return
		}

		pass.Reportf(expr.Pos(), "use AND operator with complement expression instead of AND-NOT operator")
	}
}

func checkAndNotAssignStmt(pass *analysis.Pass, stmt *ast.AssignStmt, mode string) {
	if mode == flagIgnore {
		return
	}

	if len(stmt.Lhs) != 1 {
		return
	}

	switch mode {
	case andNOTsANDNOT:
		if stmt.Tok != token.AND_ASSIGN {
			return
		}

		if astcast.ToUnaryExpr(stmt.Rhs[0]).Op != token.XOR {
			return
		}

		pass.Reportf(stmt.Pos(), "use AND-NOT assignment instead of AND assignment with complement expression")

	case andNOTsANDComp:
		if stmt.Tok != token.AND_NOT_ASSIGN {
			return
		}

		pass.Reportf(stmt.Pos(), "use AND assignment with complement expression instead of AND-NOT assignment")
	}
}
