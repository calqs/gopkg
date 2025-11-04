package handler

import (
	"net/http"

	"github.com/calqs/gopkg/router/request"
	"github.com/calqs/gopkg/router/response"
)

// Handler our basic generic request/response handler
type Handler func(*http.Request) response.Response
type (
	None     = struct{}
	NoParams = None
	N        = None
	Void     = None
)

type Request[ParamsT any] struct {
	Request *http.Request
	Params  *ParamsT
}

type Method string

func (m Method) String() string {
	return string(m)
}

type MethodHandler struct {
	Method  Method
	Handler Handler
}

type HandlerTransformer[RequestT any] interface {
	Transform() Handler
}

// GenHandler is mostly used to wrap a Handler in a generic way.
// This way, we can have handlers having a concrete struct as return type
// instead of an interface
type GenHandler[RequestT any, ResponseT response.Response] func(*Request[RequestT]) ResponseT

func (genHandler GenHandler[RequestT, ResponseT]) Transform() Handler {
	return func(r *http.Request) response.Response {
		res, err := request.ExtractData[RequestT](r)
		if err != nil {
			return response.InternalServerError("invalid data type", err)
		}
		return genHandler(&Request[RequestT]{
			Request: r,
			Params:  res,
		})
	}
}
