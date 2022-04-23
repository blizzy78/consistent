package params

type foo struct{}

func funcNoParams() {
}

func funcSingleParam(_ int) {
}

func funcExplicit(_ int, _ int) {
}

func funcExplicitNonConsecutive(_ int, _ string, _ int) {
}

func funcCompact(_, _ int) {
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

func funcLitExplicitNonConsecutive() {
	_ = func(_ int, _ string, _ int) {
	}
}

func funcLitCompact() {
	_ = func(_, _ int) {
	}
}

func (f foo) methodNoParams() {
}

func (f foo) methodSingleParam(_ int) {
}

func (f foo) methodExplicit(_ int, _ int) {
}

func (f foo) methodExplicitNonConsecutive(_ int, _ string, _ int) {
}

func (f foo) methodCompact(_, _ int) {
}
