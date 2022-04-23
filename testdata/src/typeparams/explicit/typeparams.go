package typeparams

func funcNoParams() {
}

func funcSingleParam[A int]() {
}

func funcExplicit[A int, B int]() {
}

func funcExplicitNonConsecutive[A int, B string, C int]() {
}

func funcCompact[A, B int]() { // want "declare the type of function type parameters explicitly"
}

func funcLitNoParams() {
	_ = func() {
	}
}
