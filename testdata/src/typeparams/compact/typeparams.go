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

type typeNoParams struct{}

type typeSingleParam[A int] struct{}

type typeExplicit[A int, B int] struct{} // want "declare the type of similar consecutive type parameters only once"

type typeExplicitNonConsecutive[A int, B string, C int] struct{}

type typeCompact[A, B int] struct{}
