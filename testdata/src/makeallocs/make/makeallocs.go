package makeallocs

type intSlice []int
type intMap map[int]int

func allocLiteral() {
	_ = []int{}    // want "call make instead of using slice literal"
	_ = intSlice{} // want "call make instead of using slice literal"

	_ = map[int]int{} // want "call make instead of using map literal"
	_ = intMap{}      // want "call make instead of using map literal"
}

func allocNonEmptyLiteral() {
	_ = []int{1}
	_ = intSlice{1}

	_ = map[int]int{1: 2}
	_ = intMap{1: 2}
}

func allocMake() {
	_ = make([]int, 0)
	_ = make([]int, 0, 0)
	_ = make(intSlice, 0)
	_ = make(intSlice, 0, 0)

	_ = make(map[int]int)
	_ = make(map[int]int, 0)
	_ = make(intMap)
	_ = make(intMap, 0)
}

func allocMakeNonZeroLen() {
	_ = make([]int, 1)

	_ = make(map[int]int, 1)
}

func allocMakeNonIntLen() {
	x := 0

	_ = make([]int, x)

	_ = make(map[int]int, x)
}

func allocMakeNonZeroCap() {
	_ = make([]int, 0, 1)
}

func allocMakeNonIntCap() {
	x := 1
	_ = make([]int, 0, x)
}

func makeRedefined() {
	_ = func() {
		make := func() {}
		make()
	}

	_ = func() {
		make := func(i int) {}
		make(123)
	}
}
