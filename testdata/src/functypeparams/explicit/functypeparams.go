package functypeparams

type funcNoParams func()

type funcSingleParam func(_ int)

type funcSingleParamUnnamed func(int) // want "use named function type parameters"

type funcExplicit func(_ int, _ int)

type funcExplicitNonConsecutive func(_ int, _ string, _ int)

type funcCompact func(_, _ int) // want "declare the type of function type parameters explicitly"

type iface interface {
	funcNoParams()
	funcSingleParam(_ int)
	funcSingleParamUnnamed(int) // want "use named function type parameters"
	funcExplicit(_ int, _ int)
	funcExplicitNonConsecutive(_ int, _ string, _ int)
	funcCompact(_, _ int) // want "declare the type of function type parameters explicitly"
}

func paramNoParams(_ func()) {}

func paramSingleParam(_ func(_ int)) {}

func paramSingleParamUnnamed(_ func(int)) {} // want "use named function type parameters"

func paramExplicit(_ func(_ int, _ int)) {}

func paramExplicitNonConsecutive(_ func(_ int, _ string, _ int)) {}

func paramCompact(_ func(_, _ int)) {} // want "declare the type of function type parameters explicitly"
