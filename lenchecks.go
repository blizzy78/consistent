package consistent

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	lenChecksEqualZero   = "equalZero"
	lenChecksCompareZero = "compareZero"
	lenChecksCompareOne  = "compareOne"
)

var lenChecksFlagAllowedValues = []string{flagIgnore, lenChecksEqualZero, lenChecksCompareZero, lenChecksCompareOne}

var opReverse = map[token.Token]token.Token{
	token.EQL: token.NEQ,
	token.NEQ: token.EQL,
	token.LSS: token.GTR,
	token.GTR: token.LSS,
	token.LEQ: token.GEQ,
	token.GEQ: token.LEQ,
}

func checkLenCheck(pass *analysis.Pass, expr *ast.BinaryExpr, mode string) { //nolint:gocognit,cyclop // it's not too bad
	if mode == flagIgnore {
		return
	}

	fun, oper, litInt, ok := lenCheckDetails(expr)
	if !ok {
		return
	}

	switch mode {
	case lenChecksEqualZero:
		if (oper == token.GTR && litInt == 0) ||
			(oper == token.GEQ && litInt == 1) ||
			(oper == token.LEQ && litInt == 0) ||
			(oper == token.LSS && litInt == 1) {
			pass.Reportf(expr.Pos(), "check if %s is (not) 0 instead", fun)
		}

	case lenChecksCompareZero:
		if (oper == token.NEQ && litInt == 0) ||
			(oper == token.GEQ && litInt == 1) ||
			(oper == token.EQL && litInt == 0) ||
			(oper == token.LSS && litInt == 1) {
			pass.Reportf(expr.Pos(), "compare %s to 0 instead", fun)
		}

	case lenChecksCompareOne:
		if (oper == token.NEQ && litInt == 0) ||
			(oper == token.GTR && litInt == 0) ||
			(oper == token.EQL && litInt == 0) ||
			(oper == token.LEQ && litInt == 0) {
			pass.Reportf(expr.Pos(), "compare %s to 1 instead", fun)
		}
	}
}

func lenCheckDetails(expr *ast.BinaryExpr) (string, token.Token, int, bool) {
	oper := expr.Op

	if _, ok := opReverse[oper]; !ok {
		return "", token.ILLEGAL, 0, false
	}

	call := astcast.ToCallExpr(expr.X)
	lit := astcast.ToBasicLit(expr.Y)

	if call == astcast.NilCallExpr && lit == astcast.NilBasicLit {
		lit = astcast.ToBasicLit(expr.X)
		call = astcast.ToCallExpr(expr.Y)
		oper = opReverse[oper]
	}

	if call == astcast.NilCallExpr && lit == astcast.NilBasicLit {
		return "", token.ILLEGAL, 0, false
	}

	fun := astcast.ToIdent(call.Fun).Name
	if fun != "len" && fun != "cap" {
		return "", token.ILLEGAL, 0, false
	}

	if len(call.Args) != 1 {
		return "", token.ILLEGAL, 0, false
	}

	litInt, ok := litInt(lit)
	if !ok {
		return "", token.ILLEGAL, 0, false
	}

	return fun, oper, litInt, true
}
