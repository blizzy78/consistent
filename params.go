package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var paramsFlagAllowedValues = []string{flagIgnore, fieldListExplicit, fieldListCompact}

func checkParamsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.Params, "function parameters", mode)
}

func checkParamsMethod(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.Params, "method parameters", mode)
}

func checkParamsFuncLit(pass *analysis.Pass, fun *ast.FuncLit, mode string) {
	checkFieldList(pass, fun.Type.Params, "function parameters", mode)
}
