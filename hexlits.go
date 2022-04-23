package consistent

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	hexLitsLower = "lower"
	hexLitsUpper = "upper"
)

var hexLitsFlagAllowedValues = []string{flagIgnore, hexLitsLower, hexLitsUpper}

func checkHexLit(pass *analysis.Pass, lit *ast.BasicLit, mode string) {
	if mode == flagIgnore {
		return
	}

	if !hexLit(lit) {
		return
	}

	val := lit.Value[2:]

	switch {
	case mode == hexLitsLower && strings.ToLower(val) != val:
		pass.Reportf(lit.Pos(), "use lowercase digits in hex literal")
	case mode == hexLitsUpper && strings.ToUpper(val) != val:
		pass.Reportf(lit.Pos(), "use uppercase digits in hex literal")
	}
}

func hexLit(lit *ast.BasicLit) bool {
	return lit.Kind == token.INT && (strings.HasPrefix(lit.Value, "0x") || strings.HasPrefix(lit.Value, "0X"))
}
