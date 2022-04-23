package consistent

import "testing"

func TestRun_NewAllocs(t *testing.T) {
	for _, mode := range newAllocsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "newallocs/"+mode, map[string]string{
				"newAllocs": mode,
			})
		})
	}
}
