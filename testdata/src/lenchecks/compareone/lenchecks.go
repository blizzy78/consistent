package lenchecks

func lenChecks() {
	x := []string{}

	_ = len(x) != 0 // want "compare len to 1 instead"
	_ = len(x) > 0  // want "compare len to 1 instead"
	_ = len(x) >= 1
	_ = 0 != len(x) // want "compare len to 1 instead"
	_ = 0 < len(x)  // want "compare len to 1 instead"
	_ = 1 <= len(x)
	_ = cap(x) != 0 // want "compare cap to 1 instead"
	_ = cap(x) > 0  // want "compare cap to 1 instead"
	_ = cap(x) >= 1

	_ = len(x) == 0 // want "compare len to 1 instead"
	_ = len(x) <= 0 // want "compare len to 1 instead"
	_ = len(x) < 1
	_ = 0 == len(x) // want "compare len to 1 instead"
	_ = 0 >= len(x) // want "compare len to 1 instead"
	_ = 1 > len(x)
	_ = cap(x) == 0 // want "compare cap to 1 instead"
	_ = cap(x) <= 0 // want "compare cap to 1 instead"
	_ = cap(x) < 1
}

func lenRedefined() {
	len := func() int {
		return 0
	}

	_ = len() != 0
}

func compareOther() {
	x := []string{}

	_ = len(x) != 2
	_ = len(x) == 2
	_ = len(x) < 2
	_ = len(x) <= 2
	_ = len(x) > 2
	_ = len(x) >= 2
}
