package router

import (
	"net/http"

	"github.com/calqs/gopkg/router/request"
	"github.com/calqs/gopkg/router/response"
)

// Handler our basic generic request/response handler
type Handler func(*http.Request) response.Response

// GenHandler is mostly used to wrap a Handler in a generic way.
// This way, we can have handlers having a concrete struct as return type
// instead of an interface
type GenHandler[RequestT any, ResponseT response.Response] func(*RequestT, *http.Request) ResponseT

func genHandlerToHandler[RequestT any, ResponseT response.Response](
	handler GenHandler[RequestT, ResponseT],
) func(r *http.Request) response.Response {
	return func(r *http.Request) response.Response {
		res, err := request.ExtractData[RequestT](r)
		if err != nil {
			return response.InternalServerError("invalid data type", err)
		}
		return handler(res, r)
	}
}

// Get is a wrapper around a generic handler, forcing the GET HTTP verb
func Get[RequestT any, ResponseT response.Response](handler GenHandler[RequestT, ResponseT]) MethodHandler {
	return MethodHandler{GET, genHandlerToHandler(handler)}
}

// Post is a wrapper around a generic handler, forcing the POST HTTP verb
func Post[RequestT any, ResponseT response.Response](handler GenHandler[RequestT, ResponseT]) MethodHandler {
	return MethodHandler{POST, genHandlerToHandler(handler)}
}

// Put is a wrapper around a generic handler, forcing the PUT HTTP verb
func Put[RequestT any, ResponseT response.Response](handler GenHandler[RequestT, ResponseT]) MethodHandler {
	return MethodHandler{PUT, genHandlerToHandler(handler)}
}

// Patch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func Patch[RequestT any, ResponseT response.Response](handler GenHandler[RequestT, ResponseT]) MethodHandler {
	return MethodHandler{PATCH, genHandlerToHandler(handler)}
}

// Delete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func Delete[RequestT any, ResponseT response.Response](handler GenHandler[RequestT, ResponseT]) MethodHandler {
	return MethodHandler{DELETE, genHandlerToHandler(handler)}
}
