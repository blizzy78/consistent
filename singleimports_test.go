package consistent

import "testing"

func TestRun_SingleImports(t *testing.T) {
	for _, mode := range singleImportsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "singleimports/"+mode, map[string]string{
				"singleImports": mode,
			})
		})
	}
}
