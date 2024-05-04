package consistent

import "testing"

func TestRun_NoLintDirective(t *testing.T) {
	runTest(t, "nolintdirective", map[string]string{
		"makeAllocs": makeAllocsLiteral,
	})
}
