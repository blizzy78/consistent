package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func checkReturnsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	if !namedFields(fun.Type.Results) {
		return
	}

	checkFieldList(pass, fun.Type.Results, "function return values", mode)
}

func checkReturnsMethod(pass *analysis.Pass, method *ast.FuncDecl, mode string) {
	if !namedFields(method.Type.Results) {
		return
	}

	checkFieldList(pass, method.Type.Results, "method return values", mode)
}

func checkReturnsFuncLit(pass *analysis.Pass, fun *ast.FuncLit, mode string) {
	if !namedFields(fun.Type.Results) {
		return
	}

	checkFieldList(pass, fun.Type.Results, "function return values", mode)
}

func namedFields(fields *ast.FieldList) bool {
	return fields != nil && len(fields.List) != 0 && len(fields.List[0].Names) != 0
}
