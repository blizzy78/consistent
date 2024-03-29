package consistent

import "testing"

func TestRun_LabelsRegexp(t *testing.T) {
	tests := map[string]string{
		"disable": "",
		"regexp":  "^[a-z][a-zA-Z0-9]*$",
	}

	for test, regexp := range tests {
		t.Run(test, func(t *testing.T) {
			runTest(t, "labelsregexp/"+test, map[string]string{
				"labelsRegexp": regexp,
			})
		})
	}
}
