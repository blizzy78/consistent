//go:build go1.18
// +build go1.18

package consistent

import "testing"

func TestRun_EmptyIfaces(t *testing.T) {
	for _, mode := range emptyIfacesFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "emptyifaces/"+mode, map[string]string{
				"emptyIfaces": mode,
			})
		})
	}
}
