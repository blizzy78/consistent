package functypeparams

type funcNoParams func()

type funcSingleParam func(_ int) // want "use unnamed function type parameters"

type funcSingleParamUnnamed func(int)

type funcExplicit func(_ int, _ int) // want "use unnamed function type parameters"

type funcExplicitNonConsecutive func(_ int, _ string, _ int) // want "use unnamed function type parameters"

type funcCompact func(_, _ int) // want "use unnamed function type parameters"

type iface interface {
	funcNoParams()
	funcSingleParam(_ int) // want "use unnamed function type parameters"
	funcSingleParamUnnamed(int)
	funcExplicit(_ int, _ int)                         // want "use unnamed function type parameters"
	funcExplicitNonConsecutive(_ int, _ string, _ int) // want "use unnamed function type parameters"
	funcCompact(_, _ int)                              // want "use unnamed function type parameters"
}

func paramNoParams(_ func()) {}

func paramSingleParam(_ func(_ int)) {} // want "use unnamed function type parameters"

func paramSingleParamUnnamed(_ func(int)) {}

func paramExplicit(_ func(_ int, _ int)) {} // want "use unnamed function type parameters"

func paramExplicitNonConsecutive(_ func(_ int, _ string, _ int)) {} // want "use unnamed function type parameters"

func paramCompact(_ func(_, _ int)) {} // want "use unnamed function type parameters"
