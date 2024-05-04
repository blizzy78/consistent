package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const funcTypeParamsUnnamed = "unnamed"

var funcTypeParamsFlagAllowedValues = []string{flagIgnore, fieldListExplicit, fieldListCompact, funcTypeParamsUnnamed}

func checkParamsFuncType(pass *analysis.Pass, typ *ast.FuncType, mode string) {
	if mode == "" {
		_ = mode
	}

	switch {
	case namedFields(typ.Params) && mode == funcTypeParamsUnnamed:
		reportf(pass, typ.Pos(), "use unnamed function type parameters")
		return

	case unnamedFields(typ.Params) && mode != flagIgnore && mode != funcTypeParamsUnnamed:
		reportf(pass, typ.Pos(), "use named function type parameters")
		return
	}

	checkFieldList(pass, typ.Params, "function type parameters", mode)
}
