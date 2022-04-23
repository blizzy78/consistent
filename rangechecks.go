package consistent

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astequal"
	"golang.org/x/tools/go/analysis"
)

const (
	rangeChecksLeft    = "left"
	rangeChecksCenter  = "center"
	rangeChecksRight   = "right"
	rangeChecksOutside = "outside"
)

var rangeChecksFlagAllowedValues = []string{flagIgnore, rangeChecksLeft, rangeChecksCenter}

var rangeCheckFlagDesc = map[string]string{
	rangeChecksLeft:   "write common term in range expression on the left",
	rangeChecksCenter: "write common term in range expression in the center",
}

// bits: AND OR LEFT_LESS LEFT_GREATER RIGHT_LESS RIGHT_GREATER X_LEFT X_CENTER X_RIGHT X_OUTSIDE
var rangeExprStyles = map[uint16]string{}

func init() {
	// x > low && x < high
	rangeExprStyles[0b10_01_10_1000] = rangeChecksLeft

	// x < low || x > high
	rangeExprStyles[0b01_10_01_1000] = rangeChecksLeft

	newStyles := map[uint16]string{}

	for oldBits := range rangeExprStyles {
		// build variants for center:
		// x > low && x < high  ->  low < x && x < high
		newBits := oldBits & 0b11_00_11_0000
		newBits |= ^oldBits & 0b00_11_00_0000
		newBits |= 0b00_00_00_0100
		newStyles[newBits] = rangeChecksCenter

		// build variants for right:
		// x > low && x < high  ->  low < x && high > x
		newBits = oldBits & 0b11_00_00_0000
		newBits |= ^oldBits & 0b00_11_11_0000
		newBits |= 0b00_00_00_0010
		newStyles[newBits] = rangeChecksRight

		// build variants for outside:
		// x > low && x < high  ->  x > low && high > x
		newBits = oldBits & 0b11_11_00_0000
		newBits |= ^oldBits & 0b00_00_11_0000
		newBits |= 0b00_00_00_0001
		newStyles[newBits] = rangeChecksOutside
	}

	for k, v := range newStyles {
		rangeExprStyles[k] = v
	}

	for k := range newStyles {
		delete(newStyles, k)
	}

	// build variants where left and right are swapped:
	// x > low && x < high  ->  x < high && x > low
	for oldBits, oldStyle := range rangeExprStyles {
		newBits := oldBits & 0b11_00_00_1111
		newBits |= oldBits & 0b00_11_00_0000 >> 2
		newBits |= oldBits & 0b00_00_11_0000 << 2
		newStyles[newBits] = oldStyle
	}

	for k, v := range newStyles {
		rangeExprStyles[k] = v
	}
}

func checkRangeCheck(pass *analysis.Pass, expr *ast.BinaryExpr, mode string) {
	if mode == flagIgnore {
		return
	}

	exprStyle := rangeExprStyle(expr)
	if exprStyle == "" || exprStyle == mode {
		return
	}

	pass.Reportf(expr.Pos(), rangeCheckFlagDesc[mode])
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
