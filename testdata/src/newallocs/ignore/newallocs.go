package newallocs

import (
	"io"
	"strings"
)

type foo struct {
	f int
}

func allocLiteral() {
	_ = &foo{}
	_ = &strings.Builder{}
}

func allocNew() {
	_ = new(foo)
	_ = new(strings.Builder)
	_ = new(io.Reader)
}

func allocNonEmptyLiteral() {
	_ = &foo{f: 123}
}

func callOther(i int) {
	callOther(123)
	strings.NewReader("")
}

func newRedefined() {
	new := func() {}
	new()
}

func unaryNotAnd() {
	_ = !true
}

func unaryNotLiteral() {
	x := 123
	_ = &x
}
