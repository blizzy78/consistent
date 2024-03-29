package functypeparams

type funcNoParams func()

type funcSingleParam func(_ int)

type funcSingleParamUnnamed func(int)

type funcExplicit func(_ int, _ int)

type funcExplicitNonConsecutive func(_ int, _ string, _ int)

type funcCompact func(_, _ int)

type iface interface {
	funcNoParams()
	funcSingleParam(_ int)
	funcSingleParamUnnamed(int)
	funcExplicit(_ int, _ int)
	funcExplicitNonConsecutive(_ int, _ string, _ int)
	funcCompact(_, _ int)
}

func paramNoParams(_ func()) {}

func paramSingleParam(_ func(_ int)) {}

func paramSingleParamUnnamed(_ func(int)) {}

func paramExplicit(_ func(_ int, _ int)) {}

func paramExplicitNonConsecutive(_ func(_ int, _ string, _ int)) {}

func paramCompact(_ func(_, _ int)) {}
