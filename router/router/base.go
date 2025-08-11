package router

import (
	"net/http"

	"github.com/calqs/gopkg/router/response"
)

// Layout is the context (in a setting sort of way) of a handler.
// It mostly holds dependencies and settings that need to get passed on
// up the execution tree.
type baseRouter[ConfigT any] struct {
	Config         *ConfigT
	ResponseWriter func(any, http.ResponseWriter)
}

func (s *baseRouter[ConfigT]) WithConfig(c *ConfigT) *baseRouter[ConfigT] {
	return &baseRouter[ConfigT]{
		ResponseWriter: s.ResponseWriter,
		Config:         c,
	}
}

// Handler our basic generic route handler
// @TODO: remove http.ResponseWriter?
type Handler[ConfigT any] func(http.ResponseWriter, *http.Request, ...*ConfigT) (any, *response.HttpError)

// // WithMethod is a geeneric wrapper around a generic handler, forcing the a HTTP verb
// func (l *baseRouter[ConfigT]) WithMethod(method Method, handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
// 	// #StephenCurrying
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		// a method matches
// 		if req.Method == method.String() {
// 			res, err := handler(w, req, l.Config)
// 			if err != nil {
// 				slog.Error("Handler error", "error", err.Error())
// 			}
// 			l.ResponseWriter(res, w)
// 			return
// 		}
// 	}
// }

// HandleGet is a wrapper around a generic handler, forcing the GET HTTP verb
func HandleGet[ConfigT any](handler Handler[ConfigT]) MethodHandler[ConfigT] {
	return MethodHandler[ConfigT]{GET, handler}
}

// HandlePost is a wrapper around a generic handler, forcing the POST HTTP verb
func HandlePost[ConfigT any](handler Handler[ConfigT]) MethodHandler[ConfigT] {
	return MethodHandler[ConfigT]{POST, handler}
}

// HandlePut is a wrapper around a generic handler, forcing the PUT HTTP verb
func HandlePut[ConfigT any](handler Handler[ConfigT]) MethodHandler[ConfigT] {
	return MethodHandler[ConfigT]{PUT, handler}
}

// HandlePatch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func HandlePatch[ConfigT any](handler Handler[ConfigT]) MethodHandler[ConfigT] {
	return MethodHandler[ConfigT]{PATCH, handler}
}

// HandleDelete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func HandleDelete[ConfigT any](handler Handler[ConfigT]) MethodHandler[ConfigT] {
	return MethodHandler[ConfigT]{DELETE, handler}
}
