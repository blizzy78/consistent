package newallocs

import "strings"

type foo struct {
	f int
}

func allocLiteral() {
	_ = &foo{}
	_ = &strings.Builder{}
}

func allocNew() {
	_ = new(foo)             // want "use zero-value literal instead of calling new"
	_ = new(strings.Builder) // want "use zero-value literal instead of calling new"
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
