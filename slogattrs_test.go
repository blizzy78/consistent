package consistent

import (
	"strings"
	"testing"
)

func TestRun_SlogAttrs(t *testing.T) {
	for _, mode := range slogAttrsFlagAllowedValues {
		t.Run(mode, func(t *testing.T) {
			runTest(t, "slogattrs/"+strings.ToLower(mode), map[string]string{
				"slogAttrs": mode,
			})
		})
	}
}
