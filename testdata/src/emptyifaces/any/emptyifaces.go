package emptyifaces

type ifaceType interface{} // want "use any instead of interface{}"

type anyType any

type otherType interface {
	foo()
}

func ifaceParam(_ interface{}) { // want "use any instead of interface{}"
}

func anyParam(_ any) {
}

func ifaceReturn() interface{} { // want "use any instead of interface{}"
	return nil
}

func anyReturn() any {
	return nil
}

func ifaceTypeParam[T interface{}](_ T) { // want "use any instead of interface{}"
}

func anyTypeParam[T any](_ T) {
}

func ifaceVar() {
	var _ interface{} // want "use any instead of interface{}"
}

func anyVar() {
	var _ any
}
