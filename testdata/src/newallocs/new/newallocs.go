package newallocs

import "strings"

type foo struct {
	f int
}

func allocLiteral() {
	_ = &foo{}             // want "call new instead of using zero-value literal"
	_ = &strings.Builder{} // want "call new instead of using zero-value literal"
}

func allocNew() {
	_ = new(foo)
	_ = new(strings.Builder)
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
