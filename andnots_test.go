package consistent

import (
	"strings"
	"testing"
)

func TestRun_AndNots(t *testing.T) {
	for _, mode := range andNOTsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "andnots/"+strings.ToLower(mode), map[string]string{
				"andNOTs": mode,
			})
		})
	}
}
