package public

import (
	"net/http"

	"github.com/calqs/gopkg/router/handler"
	"github.com/calqs/gopkg/router/response"
)

// Get is a wrapper around a generic handler, forcing the GET HTTP verb
func Get[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) handler.MethodHandler {
	return handler.MethodHandler{Method: http.MethodGet, Handler: handler.GenHandlerToHandler(h)}
}

// Post is a wrapper around a generic handler, forcing the POST HTTP verb
func Post[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) handler.MethodHandler {
	return handler.MethodHandler{Method: http.MethodPost, Handler: handler.GenHandlerToHandler(h)}
}

// Put is a wrapper around a generic handler, forcing the PUT HTTP verb
func Put[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) handler.MethodHandler {
	return handler.MethodHandler{Method: http.MethodPut, Handler: handler.GenHandlerToHandler(h)}
}

// Patch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func Patch[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) handler.MethodHandler {
	return handler.MethodHandler{Method: http.MethodPatch, Handler: handler.GenHandlerToHandler(h)}
}

// Delete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func Delete[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) handler.MethodHandler {
	return handler.MethodHandler{Method: http.MethodDelete, Handler: handler.GenHandlerToHandler(h)}
}
