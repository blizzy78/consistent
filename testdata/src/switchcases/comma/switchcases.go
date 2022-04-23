package switchcases

func switchCases() {
	switch {
	case 1 < 2:
	case 1 < 2, 2 < 3:
	case 1 < 2, 2 < 3, 3 < 4:
	case 1 < 2 || 2 < 3: // want "separate cases with comma instead of using logical OR"
	case 1 < 2 || 2 < 3 || 3 < 4: // want "separate cases with comma instead of using logical OR"
	}
}

func tag() {
	x := 1
	switch x {
	case 1:
	case 2, 3:
	}
}

func defaultCase() {
	switch {
	default:
	}
}

func notOR() {
	switch {
	case 1 < 2 && 2 < 3:
	case 1 < 2 || 2 < 3 && 3 < 4:
	case 1 < 2 && 2 < 3 || 3 < 4:
	}
}
