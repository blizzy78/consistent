package consistent

import "testing"

func TestRun_Params(t *testing.T) {
	for _, mode := range fieldListFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "params/"+mode, map[string]string{
				"params": mode,
			})
		})
	}
}
