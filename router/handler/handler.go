package handler

import (
	"net/http"

	"github.com/calqs/gopkg/router/request"
	"github.com/calqs/gopkg/router/response"
)

// Handler our basic generic request/response handler
type Handler func(*http.Request) response.Response

type Request[ParamsT any] struct {
	Request *http.Request
	Params  *ParamsT
}

// GenHandler is mostly used to wrap a Handler in a generic way.
// This way, we can have handlers having a concrete struct as return type
// instead of an interface
type GenHandler[RequestT any, ResponseT response.Response] func(*Request[RequestT]) ResponseT

type Method string

func (m Method) String() string {
	return string(m)
}

type MethodHandler struct {
	Method  Method
	Handler Handler
}

func GenHandlerToHandler[RequestT any, ResponseT response.Response](
	handler GenHandler[RequestT, ResponseT],
) func(r *http.Request) response.Response {
	return func(r *http.Request) response.Response {
		res, err := request.ExtractData[RequestT](r)
		if err != nil {
			return response.InternalServerError("invalid data type", err)
		}
		return handler(&Request[RequestT]{
			Request: r,
			Params:  res,
		})
	}
}
