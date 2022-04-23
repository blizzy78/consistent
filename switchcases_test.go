package consistent

import (
	"strings"
	"testing"
)

func TestRun_SwitchCases(t *testing.T) {
	for _, mode := range switchCasesFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "switchcases/"+strings.ToLower(mode), map[string]string{
				"switchCases": mode,
				"rangeChecks": "ignore",
			})
		})
	}
}
