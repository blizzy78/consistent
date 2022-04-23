package andnots

func andNot() {
	x := 123

	_ = 123 &^ 234 // want "use AND operator with complement expression instead of AND-NOT operator"
	x &^= 234      // want "use AND assignment with complement expression instead of AND-NOT assignment"

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
