package typeparams

func funcNoParams() {
}

func funcSingleParam[A int]() {
}

func funcExplicit[A int, B int]() { // want "declare the type of similar consecutive function type parameters only once"
}

func funcExplicitNonConsecutive[A int, B string, C int]() {
}

func funcCompact[A, B int]() {
}

func funcLitNoParams() {
	_ = func() {
	}
}
