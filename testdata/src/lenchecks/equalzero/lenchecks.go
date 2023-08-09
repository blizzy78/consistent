package lenchecks

func lenChecks() {
	x := []string{}

	_ = len(x) != 0
	_ = len(x) > 0  // want `check if len is not 0 instead`
	_ = len(x) >= 1 // want `check if len is not 0 instead`
	_ = 0 != len(x)
	_ = 0 < len(x)  // want `check if len is not 0 instead`
	_ = 1 <= len(x) // want `check if len is not 0 instead`
	_ = cap(x) != 0
	_ = cap(x) > 0  // want `check if cap is not 0 instead`
	_ = cap(x) >= 1 // want `check if cap is not 0 instead`

	_ = len(x) == 0
	_ = len(x) <= 0 // want `check if len is 0 instead`
	_ = len(x) < 1  // want `check if len is 0 instead`
	_ = 0 == len(x)
	_ = 0 >= len(x) // want `check if len is 0 instead`
	_ = 1 > len(x)  // want `check if len is 0 instead`
	_ = cap(x) == 0
	_ = cap(x) <= 0 // want `check if cap is 0 instead`
	_ = cap(x) < 1  // want `check if cap is 0 instead`
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
