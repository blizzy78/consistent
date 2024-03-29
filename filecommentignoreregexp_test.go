package consistent

import "testing"

func TestRun_FileCommentIgnoreRegexp(t *testing.T) {
	tests := map[string]string{
		// "disable": "",
		"regexp": "ignore this file: yes",
	}

	for test, regexp := range tests {
		t.Run(test, func(t *testing.T) {
			runTest(t, "filecommentignoreregexp/"+test, map[string]string{
				"fileCommentIgnoreRegexp": regexp,
			})
		})
	}
}
