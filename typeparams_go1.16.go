//go:build (go1.16 && !go1.18) || (go1.17 && !go1.18)
// +build go1.16,!go1.18 go1.17,!go1.18

package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func checkTypeParamsFunc(pass *analysis.Pass, fun *ast.FuncDecl, mode string) {
}
