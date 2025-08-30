package public

import (
	"net/http"

	"github.com/calqs/gopkg/router/handler"
	"github.com/calqs/gopkg/router/response"
)

func routeIt[RequestT any, ResponseT response.Response](method string, h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			return
		}
		res := handler.GenHandlerToHandler(h)(r)
		res.Send(w)
	})
}

// Get is a wrapper around a generic handler, forcing the GET HTTP verb
func Get[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return routeIt(http.MethodGet, h)
}

// Post is a wrapper around a generic handler, forcing the POST HTTP verb
func Post[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return routeIt(http.MethodPost, h)
}

// Put is a wrapper around a generic handler, forcing the PUT HTTP verb
func Put[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return routeIt(http.MethodPut, h)
}

// Patch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func Patch[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return routeIt(http.MethodPatch, h)
}

// Delete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func Delete[RequestT any, ResponseT response.Response](h handler.GenHandler[RequestT, ResponseT]) http.Handler {
	return routeIt(http.MethodDelete, h)
}
