package consistent

import (
	"go/ast"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
)

const (
	switchDefaultsLast  = "last"
	switchDefaultsFirst = "first"
)

var switchDefaultsFlagAllowedValues = []string{flagIgnore, switchDefaultsLast, switchDefaultsFirst}

func checkSwitchDefault(pass *analysis.Pass, stmt *ast.SwitchStmt, mode string) {
	if mode == flagIgnore {
		return
	}

	idx := defaultClauseIndex(stmt)
	if idx < 0 {
		return
	}

	switch mode {
	case switchDefaultsLast:
		if idx == len(stmt.Body.List)-1 {
			return
		}

		reportf(pass, stmt.Body.List[idx].Pos(), "move switch default clause to the end")

	case switchDefaultsFirst:
		if idx == 0 {
			return
		}

		reportf(pass, stmt.Body.List[idx].Pos(), "move switch default clause to the beginning")
	}
}

func defaultClauseIndex(stmt *ast.SwitchStmt) int {
	for i, c := range stmt.Body.List {
		if astcast.ToCaseClause(c).List == nil {
			return i
		}
	}

	return -1
}
