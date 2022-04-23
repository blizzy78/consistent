package consistent

import (
	"strings"
	"testing"
)

func TestRun_SwitchDefaults(t *testing.T) {
	for _, mode := range switchDefaultsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "switchdefaults/"+strings.ToLower(mode), map[string]string{
				"switchDefaults": mode,
			})
		})
	}
}
