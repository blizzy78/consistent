package consistent

import "testing"

func TestRun_Returns(t *testing.T) {
	for _, mode := range fieldListFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "returns/"+mode, map[string]string{
				"returns": mode,
			})
		})
	}
}
