package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var typeParamsFlagAllowedValues = []string{flagIgnore, fieldListExplicit, fieldListCompact}

func checkTypeParamsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.TypeParams, "function type parameters", mode)
}

func checkTypeParamsType(pass *analysis.Pass, spec *ast.TypeSpec, mode string) {
	checkFieldList(pass, spec.TypeParams, "type parameters", mode)
}
