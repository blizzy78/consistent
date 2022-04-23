package switchdefaults

func switchDefault() {
	switch {
	case 1 < 2:
	default: // want "move switch default clause to the beginning"
	}

	switch {
	default:
	case 1 < 2:
	}

	switch {
	case 1 < 2:
	default: // want "move switch default clause to the beginning"
	case 3 < 4:
	}
}

func noDefault() {
	switch {
	}

	switch {
	case 1 < 2:
	}
}

func defaultOnly() {
	switch {
	default:
	}
}
