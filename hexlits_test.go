package consistent

import "testing"

func TestRun_HexLits(t *testing.T) {
	for _, mode := range hexLitsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "hexlits/"+mode, map[string]string{
				"hexLits": mode,
			})
		})
	}
}
