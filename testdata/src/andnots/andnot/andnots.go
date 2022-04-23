package andnots

func andNot() {
	x := 123

	_ = 123 &^ 234
	x &^= 234

	_ = 123 & ^234 // want "use AND-NOT operator instead of AND operator with complement expression"
	x &= ^234      // want "use AND-NOT assignment instead of AND assignment with complement expression"
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
