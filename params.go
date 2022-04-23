package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func checkParamsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.Params, "function arguments", mode)
}

func checkParamsMethod(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.Params, "method arguments", mode)
}

func checkParamsFuncLit(pass *analysis.Pass, fun *ast.FuncLit, mode string) {
	checkFieldList(pass, fun.Type.Params, "function arguments", mode)
}
