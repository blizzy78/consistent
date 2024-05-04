package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const (
	fieldListExplicit = "explicit"
	fieldListCompact  = "compact"
)

func checkFieldList(pass *analysis.Pass, fields *ast.FieldList, fieldTypePlural string, mode string) {
	if fields == nil {
		return
	}

	switch mode {
	case flagIgnore:

	case fieldListExplicit:
		for _, f := range fields.List {
			if len(f.Names) > 1 {
				reportf(pass, fields.Pos(), "declare the type of %s explicitly", fieldTypePlural)
				break
			}
		}

	case fieldListCompact:
		list := fields.List
		if len(list) <= 1 {
			return
		}

		for i, f := range list[1:] {
			typ := pass.TypesInfo.TypeOf(f.Type)
			prevTyp := pass.TypesInfo.TypeOf(list[i].Type)

			if typ == prevTyp {
				reportf(pass, fields.Pos(), "declare the type of similar consecutive %s only once", fieldTypePlural)
				break
			}
		}
	}
}
