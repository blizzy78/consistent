package consistent

import "testing"

func TestRun_TypeParams(t *testing.T) {
	for _, mode := range fieldListFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "typeparams/"+mode, map[string]string{
				"typeParams": mode,
			})
		})
	}
}
