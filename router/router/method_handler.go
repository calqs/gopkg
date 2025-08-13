package router

type Method string

func (m Method) String() string {
	return string(m)
}

type MethodHandler struct {
	Method  Method
	Handler Handler
}
