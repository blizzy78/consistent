package nolintdirective

func foo() {
	_ = make([]int, 0) // want "use slice literal instead of calling make"
	_ = make([]int, 0) //nolint:consistent
	_ = make([]int, 0) //nolint:consistent // ignore
	_ = make([]int, 0) //nolint:foo,consistent
	_ = make([]int, 0) //nolint:foo,consistent // ignore
	_ = make([]int, 0) //nolint:consistent,foo
	_ = make([]int, 0) //nolint:consistent,foo // ignore
	_ = make([]int, 0) //nolint:foo,consistent,bar
	_ = make([]int, 0) //nolint:foo,consistent,bar // ignore
	_ = make([]int, 0) //nolint: foo, consistent, bar
	_ = make([]int, 0) //nolint: foo, consistent, bar // ignore
}
