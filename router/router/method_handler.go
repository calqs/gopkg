package router

type Method string

func (m Method) String() string {
	return string(m)
}

type MethodHandler[ConfigT any] struct {
	Method  Method
	Handler Handler[ConfigT]
}
