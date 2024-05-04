package consistent

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	switchCasesComma = "comma"
	switchCasesOR    = "or"
)

var switchCasesFlagAllowedValues = []string{flagIgnore, switchCasesComma, switchCasesOR}

func checkSwitchCases(pass *analysis.Pass, stmt *ast.SwitchStmt, mode string) {
	if mode == flagIgnore {
		return
	}

	if stmt.Tag != nil {
		return
	}

	for _, clause := range stmt.Body.List {
		checkSwitchCase(pass, astcast.ToCaseClause(clause), mode)
	}
}

func checkSwitchCase(pass *analysis.Pass, clause *ast.CaseClause, mode string) {
	switch mode {
	case switchCasesComma:
		for _, expr := range clause.List {
			log, ok := toLogical(expr)
			if !ok {
				continue
			}

			if !isLogicalORRecursive(log) {
				continue
			}

			reportf(pass, expr.Pos(), "separate cases with comma instead of using logical OR")
		}

	case switchCasesOR:
		if len(clause.List) <= 1 {
			return
		}

		reportf(pass, clause.Pos(), "use logical OR instead of separating cases with comma")
	}
}

func toLogical(expr ast.Expr) (*ast.BinaryExpr, bool) {
	bin := astcast.ToBinaryExpr(expr)
	return bin, bin.Op == token.LAND || bin.Op == token.LOR
}

func isLogicalORRecursive(expr *ast.BinaryExpr) bool {
	if expr.Op != token.LOR {
		return false
	}

	if log, ok := toLogical(expr.X); ok && !isLogicalORRecursive(log) {
		return false
	}

	if log, ok := toLogical(expr.Y); ok && !isLogicalORRecursive(log) {
		return false
	}

	return true
}
