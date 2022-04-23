package returns

type foo struct{}

func funcNoReturn() {
}

func funcReturnAnonSingle() int {
	return 0
}

func funcReturnAnon() (int, int) {
	return 0, 0
}

func funcExplicit() (a int, b int) { // want "declare the type of similar consecutive function return values only once"
	return 0, 0
}

func funcExplicitNonConsecutive() (a int, b string, c int) {
	return 0, "", 0
}

func funcCompact() (a, b int) {
	return 0, 0
}

func funcLitNoReturn() {
	_ = func() {
	}
}

func funcLitReturnAnonSingle() {
	_ = func() int {
		return 0
	}
}

func funcLitReturnAnon() {
	_ = func() (int, int) {
		return 0, 0
	}
}

func funcLitExplicit() {
	_ = func() (a int, b int) { // want "declare the type of similar consecutive function return values only once"
		return 0, 0
	}
}

func funcLitExplicitNonConsecutive() {
	_ = func() (a int, b string, c int) {
		return 0, "", 0
	}
}

func funcLitCompact() {
	_ = func() (a, b int) {
		return 0, 0
	}
}

func (f foo) methodNoReturn() {
}

func (f foo) methodReturnAnonSingle() int {
	return 0
}

func (f foo) methodReturnAnon() (int, int) {
	return 0, 0
}

func (f foo) methodExplicit() (a int, b int) { // want "declare the type of similar consecutive method return values only once"
	return 0, 0
}

func (f foo) methodExplicitNonConsecutive() (a int, b string, c int) {
	return 0, "", 0
}

func (f foo) methodCompact() (a, b int) {
	return 0, 0
}
