//go:build go1.18
// +build go1.18

package consistent

import (
	"go/ast"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

func checkEmptyIface(pass *analysis.Pass, node ast.Node, mode string) {
	switch mode {
	case flagIgnore:
		return

	case emptyIfacesAny:
		checkEmptyIfaceAny(pass, node)

	case emptyIfacesIface:
		checkEmptyIfaceIface(pass, node)
	}
}

func checkEmptyIfaceAny(pass *analysis.Pass, node ast.Node) {
	itype := astcast.ToInterfaceType(node)
	if itype == astcast.NilInterfaceType {
		return
	}

	if itype.Methods != nil && itype.Methods.List != nil && len(itype.Methods.List) != 0 {
		return
	}

	pass.Reportf(node.Pos(), "use any instead of interface{}")
}

func checkEmptyIfaceIface(pass *analysis.Pass, node ast.Node) {
	var typ ast.Expr

	switch node2 := node.(type) {
	case *ast.Field:
		typ = node2.Type
	case *ast.TypeSpec:
		typ = node2.Type
	case *ast.ValueSpec:
		typ = node2.Type
	}

	if ident := astcast.ToIdent(typ); ident.Name != "any" {
		return
	}

	pass.Reportf(typ.Pos(), "use interface{} instead of any")
}
