package consistent

import "testing"

func TestRun_RangeChecks(t *testing.T) {
	for _, mode := range rangeChecksFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "rangechecks/"+mode, map[string]string{
				"rangeChecks": mode,
			})
		})
	}
}
