package consistent

import "testing"

func TestRun_MakeAllocs(t *testing.T) {
	for _, mode := range makeAllocsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "makeallocs/"+mode, map[string]string{
				"makeAllocs": mode,
			})
		})
	}
}
