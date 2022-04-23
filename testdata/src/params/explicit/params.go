package params

type foo struct{}

func funcNoParams() {
}

func funcSingleParam(_ int) {
}

func funcExplicit(_ int, _ int) {
}

func funcExplicitNotConsecutive(_ int, _ string, _ int) {
}

func funcCompact(_, _ int) { // want "declare the type of function arguments explicitly"
}

func funcLitNoParams() {
	_ = func() {
	}
}

func funcLitSingleParam() {
	_ = func(_ int) {
	}
}

func funcLitExplicit() {
	_ = func(_ int, _ int) {
	}
}

func funcLitExplicitNotConsecutive() {
	_ = func(_ int, _ string, _ int) {
	}
}

func funcLitCompact() {
	_ = func(_, _ int) { // want "declare the type of function arguments explicitly"
	}
}

func (f foo) methodNoParams() {
}

func (f foo) methodSingleParam(_ int) {
}

func (f foo) methodExplicit(_ int, _ int) {
}

func (f foo) methodExplicitNotConsecutive(_ int, _ string, _ int) {
}

func (f foo) methodCompact(_, _ int) { // want "declare the type of method arguments explicitly"
}
