package consistent

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	floatLitsExplicit = "explicit"
	floatLitsImplicit = "implicit"
)

var floatLitsFlagAllowedValues = []string{flagIgnore, floatLitsExplicit, floatLitsImplicit}

func checkFloatLit(pass *analysis.Pass, lit *ast.BasicLit, mode string) {
	if mode == flagIgnore {
		return
	}

	if lit.Kind != token.FLOAT {
		return
	}

	value := strings.TrimPrefix(lit.Value, "-")
	dotPos := strings.Index(value, ".")

	if dotPos < 0 {
		return
	}

	switch implicit := dotPos == 0; {
	case mode == floatLitsExplicit && implicit:
		pass.Reportf(lit.Pos(), "add zero before decimal point in floating-point literal")
	case mode == floatLitsImplicit && !implicit && strings.TrimLeft(value[:dotPos], "0") == "":
		pass.Reportf(lit.Pos(), "remove zero before decimal point in floating-point literal")
	}
}
