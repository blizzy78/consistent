package consistent

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func checkLabel(pass *analysis.Pass, stmt *ast.LabeledStmt, regexp regexpValue) {
	if regexp.r == nil {
		return
	}

	if regexp.r.MatchString(stmt.Label.Name) {
		return
	}

	pass.Reportf(stmt.Label.Pos(), "change label to match regular expression: %s", regexp.s)
}
