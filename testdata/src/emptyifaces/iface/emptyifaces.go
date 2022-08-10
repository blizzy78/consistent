package emptyifaces

type ifaceType interface{}

type anyType any // want "use interface{} instead of any"

type otherType interface {
	foo()
}

func ifaceParam(_ interface{}) {
}

func anyParam(_ any) { // want "use interface{} instead of any"
}

func ifaceReturn() interface{} {
	return nil
}

func anyReturn() any { // want "use interface{} instead of any"
	return nil
}

func ifaceTypeParam[T interface{}](_ T) {
}

func anyTypeParam[T any](_ T) { // want "use interface{} instead of any"
}

func ifaceVar() {
	var _ interface{}
}

func anyVar() {
	var _ any // want "use interface{} instead of any"
}
