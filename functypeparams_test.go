package consistent

import "testing"

func TestRun_FuncTypeParams(t *testing.T) {
	for _, mode := range funcTypeParamsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "functypeparams/"+mode, map[string]string{
				"funcTypeParams": mode,
			})
		})
	}
}
