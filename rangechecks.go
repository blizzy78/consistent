package consistent

//go:generate go run ./cmd/rangeexprstyles rangechecks_styles.go
//go:generate gofmt -w rangechecks_styles.go

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astequal"
	"golang.org/x/tools/go/analysis"
)

const (
	rangeChecksLeft   = "left"
	rangeChecksCenter = "center"
)

var rangeChecksFlagAllowedValues = []string{flagIgnore, rangeChecksLeft, rangeChecksCenter}

func checkRangeCheck(pass *analysis.Pass, expr *ast.BinaryExpr, mode string) {
	if mode == flagIgnore {
		return
	}

	exprStyle := rangeExprStyle(expr)
	if exprStyle == "" || exprStyle == mode {
		return
	}

	switch mode {
	case rangeChecksLeft:
		reportf(pass, expr.Pos(), "write common term in range expression on the left")
	case rangeChecksCenter:
		reportf(pass, expr.Pos(), "write common term in range expression in the center")
	}
}

func rangeExprStyle(expr *ast.BinaryExpr) string { //nolint:cyclop,gocognit // collecting a bunch of flags
	if expr.Op != token.LAND && expr.Op != token.LOR {
		return ""
	}

	left := astcast.ToBinaryExpr(expr.X)
	if !isCompareExpr(left) {
		return ""
	}

	right := astcast.ToBinaryExpr(expr.Y)
	if !isCompareExpr(right) {
		return ""
	}

	exprBits := uint16(0)

	if expr.Op == token.LAND {
		exprBits |= 1
	}

	exprBits <<= 1

	if expr.Op == token.LOR {
		exprBits |= 1
	}

	exprBits <<= 1

	if left.Op == token.LSS || left.Op == token.LEQ {
		exprBits |= 1
	}

	exprBits <<= 1

	if left.Op == token.GTR || left.Op == token.GEQ {
		exprBits |= 1
	}

	exprBits <<= 1

	if right.Op == token.LSS || right.Op == token.LEQ {
		exprBits |= 1
	}

	exprBits <<= 1

	if right.Op == token.GTR || right.Op == token.GEQ {
		exprBits |= 1
	}

	exprBits <<= 1

	if astequal.Expr(left.X, right.X) {
		exprBits |= 1
	}

	exprBits <<= 1

	if astequal.Expr(left.Y, right.X) {
		exprBits |= 1
	}

	exprBits <<= 1

	if astequal.Expr(left.Y, right.Y) {
		exprBits |= 1
	}

	exprBits <<= 1

	if astequal.Expr(left.X, right.Y) {
		exprBits |= 1
	}

	// end of bits, no shift here

	return rangeExprStyles[exprBits]
}

func isCompareExpr(expr *ast.BinaryExpr) bool {
	switch expr.Op {
	case token.LSS, token.LEQ, token.GTR, token.GEQ:
		return true
	default:
		return false
	}
}
