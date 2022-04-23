package labelsregexp

func labelsRegexp() {
label:
	for {
		break label
	}

labelFoo2:
	for {
		break labelFoo2
	}

label_3: // want "change label to match regular expression: .*"
	for {
		break label_3
	}
}
