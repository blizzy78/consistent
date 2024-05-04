package consistent

import (
	"go/ast"
	"go/types"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	slogAttrsBare       = "bare"
	slogAttrsAttr       = "attr"
	slogAttrsConsistent = "consistent"
)

type slogFunc struct {
	name     string
	skipArgs int
}

var slogAttrsFlagAllowedValues = []string{flagIgnore, slogAttrsBare, slogAttrsAttr, slogAttrsConsistent}

var (
	slogLoggerFuncs = []*slogFunc{
		{name: "Debug", skipArgs: 1},
		{name: "DebugContext", skipArgs: 2},
		{name: "Error", skipArgs: 1},
		{name: "ErrorContext", skipArgs: 2},
		{name: "Info", skipArgs: 1},
		{name: "InfoContext", skipArgs: 2},
		{name: "Log", skipArgs: 3},
		{name: "Warn", skipArgs: 1},
		{name: "WarnContext", skipArgs: 2},
		{name: "With"},
	}

	slogPackageFuncs = append(
		append([]*slogFunc{}, slogLoggerFuncs...),
		&slogFunc{name: "Group", skipArgs: 1},
	)
)

func checkSlogAttrs(pass *analysis.Pass, call *ast.CallExpr, mode string) {
	if mode == flagIgnore {
		return
	}

	bare, attrs := slogCallArgs(pass, call)

	switch {
	case mode == slogAttrsBare && attrs:
		reportf(pass, call.Pos(), "use bare arguments only")

	case mode == slogAttrsAttr && bare:
		reportf(pass, call.Pos(), "use Attr arguments only")

	case mode == slogAttrsConsistent && bare && attrs:
		reportf(pass, call.Pos(), "use consistent arguments (either bare or Attr)")
	}
}

func slogCallArgs(pass *analysis.Pass, call *ast.CallExpr) (bool, bool) {
	fun := slogFunction(pass, call)
	if fun == nil {
		return false, false
	}

	attrTyp := importPackageLookupType(pass, "log/slog", "Attr")
	if attrTyp == nil {
		return false, false
	}

	bare := false
	attrs := false

	args := call.Args[fun.skipArgs:]

	for idx := 0; idx < len(args); idx++ {
		expr := args[idx]
		exprTyp := pass.TypesInfo.TypeOf(expr)

		if types.AssignableTo(exprTyp, attrTyp) {
			attrs = true
		} else {
			bare = true

			// skip value
			idx++
		}
	}

	return bare, attrs
}

func slogFunction(pass *analysis.Pass, call *ast.CallExpr) *slogFunc {
	sel := astcast.ToSelectorExpr(call.Fun)
	if sel == astcast.NilSelectorExpr {
		return nil
	}

	if fun := slogPackageFunction(sel); fun != nil {
		return fun
	}

	return slogLoggerFunction(pass, sel)
}

func slogPackageFunction(sel *ast.SelectorExpr) *slogFunc {
	pkg := astcast.ToIdent(sel.X).Name
	if pkg != "slog" {
		return nil
	}

	funIdent := astcast.ToIdent(sel.Sel)
	if funIdent == astcast.NilIdent {
		return nil
	}

	for _, fun := range slogPackageFuncs {
		if fun.name == funIdent.Name {
			return fun
		}
	}

	return nil
}

func slogLoggerFunction(pass *analysis.Pass, sel *ast.SelectorExpr) *slogFunc {
	slogLoggerType := importPackageLookupType(pass, "log/slog", "Logger")
	if slogLoggerType == nil {
		return nil
	}

	slogLoggerType = types.NewPointer(slogLoggerType)

	loggerTyp := pass.TypesInfo.TypeOf(sel.X)

	if !types.AssignableTo(loggerTyp, slogLoggerType) {
		return nil
	}

	fun := astcast.ToIdent(sel.Sel).Name
	for _, loggerFun := range slogLoggerFuncs {
		if loggerFun.name == fun {
			return loggerFun
		}
	}

	return nil
}

func importPackageLookupType(pass *analysis.Pass, pkgName string, name string) types.Type {
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() != pkgName {
			continue
		}

		return pkg.Scope().Lookup(name).Type()
	}

	return nil
}
