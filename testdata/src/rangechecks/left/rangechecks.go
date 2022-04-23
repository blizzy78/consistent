package rangechecks

func rangeChecks() {
	low := 1
	high := 10
	x := 5

	_ = x > low && x < high
	_ = low < x && x < high // want "write common term in range expression on the left"
	_ = low < x && high > x // want "write common term in range expression on the left"
	_ = x > low && high > x // want "write common term in range expression on the left"
	_ = x*2 > low && x*2 < high
	_ = len("foo") > low && len("foo") < high

	_ = x < low || x > high
	_ = low > x || x > high // want "write common term in range expression on the left"
	_ = low > x || high < x // want "write common term in range expression on the left"
	_ = x < low || high < x // want "write common term in range expression on the left"
	_ = x*2 < low || x*2 > high
	_ = len("foo") < low || len("foo") > high
}

func noCommon() {
	low := 1
	high := 10
	x := 5
	y := 6

	_ = x > low && y < high
}

func nonIdent() {
	low := 1
	high := 10
	x := 5

	_ = x > low && "a" < "b"
	_ = "a" < "b" && x < high
}

func nonANDOrOR() {
	_ = 0 | 1
}

func nonBinary() {
	_ = true && 0 < 1
	_ = 0 < 1 && true
}

func nonCompare() {
	_ = 0 == 1 && 0 < 1
	_ = 0 < 1 && 0 == 1
}
