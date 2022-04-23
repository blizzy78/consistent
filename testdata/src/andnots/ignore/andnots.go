package andnots

func andNot() {
	x := 123

	_ = 123 &^ 234
	x &^= 234

	_ = 123 & ^234
	x &= ^234
}

func nonUnary() {
	_ = 123 & 234
}

func nonXOR() {
	x := 123

	_ = 123 & -234
	x &= -234
}

func multi() {
	_, _ = func() (int, int) {
		return 1, 2
	}()
}
