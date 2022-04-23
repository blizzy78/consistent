package consistent

import (
	"strings"
	"testing"
)

func TestRun_LenChecks(t *testing.T) {
	for _, mode := range lenChecksFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "lenchecks/"+strings.ToLower(mode), map[string]string{
				"lenChecks": mode,
			})
		})
	}
}
