package switchdefaults

func switchDefault() {
	switch {
	case 1 < 2:
	default:
	}

	switch {
	default:
	case 1 < 2:
	}
}

func noDefault() {
	switch {
	}

	switch {
	case 1 < 2:
	}
}

func inBetween() {
	switch {
	case 1 < 2:
	default:
	case 3 < 4:
	}
}

func defaultOnly() {
	switch {
	default:
	}
}
