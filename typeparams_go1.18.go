//go:build go1.18
// +build go1.18

package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func checkTypeParamsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
	checkFieldList(pass, fun.Type.TypeParams, "function type parameters", mode)
}
