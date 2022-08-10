package emptyifaces

type ifaceType interface{}

type anyType any

type otherType interface {
	foo()
}

func ifaceParam(_ interface{}) {
}

func anyParam(_ any) {
}

func ifaceReturn() interface{} {
	return nil
}

func anyReturn() any {
	return nil
}

func ifaceTypeParam[T interface{}](_ T) {
}

func anyTypeParam[T any](_ T) {
}

func ifaceVar() {
	var _ interface{}
}

func anyVar() {
	var _ any
}
