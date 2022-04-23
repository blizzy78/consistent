package consistent

import "testing"

func TestRun_FloatLits(t *testing.T) {
	for _, mode := range floatLitsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "floatlits/"+mode, map[string]string{
				"floatLits": mode,
			})
		})
	}
}
