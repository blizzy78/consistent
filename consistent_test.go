package consistent

import (
	"os"
	"testing"

	"github.com/go-toolsmith/strparse"
	"github.com/matryer/is"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLitInt(t *testing.T) {
	tests := []struct {
		given string
		want  int
	}{
		{"0", 0},
		{"123", 123},
		{"1_2_3", 123},
		{"0x123", 0x123},
		{"0X123", 0x123},
		{"0b11011", 0b11011},
		{"0B11011", 0b11011},
		{"0644", 0644},
		{"0o644", 0644},
		{"0O644", 0644},
		{"000644", 0644},
	}

	for _, test := range tests {
		t.Run(test.given, func(t *testing.T) {
			is := is.New(t)
			i, ok := litInt(strparse.Expr(test.given))
			is.Equal(i, test.want)
			is.True(ok)
		})
	}
}

func runTest(t *testing.T, pkg string, flags map[string]string) {
	t.Helper()

	analyzer := NewAnalyzer()

	for k, v := range flags {
		_ = analyzer.Flags.Set(k, v)
	}

	wd, _ := os.Getwd()
	analysistest.Run(t, wd+"/testdata", analyzer, pkg)
}
