package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/calqs/http/response"
)

// Layout is the context (in a setting sort of way) of a handler.
// It mostly holds dependencies and settings that need to get passed on
// up the execution tree.
type Server[ConfigT any] struct {
	Config         *ConfigT
	ResponseWriter func(*any, http.ResponseWriter)
}

// Handler our basic generic route handler
type Handler[ConfigT any] func(http.ResponseWriter, *http.Request, ...*ConfigT) (*any, *response.HttpError)

// Methods vector of available HTTP MEthods
var Methods = [5]string{"GET", "POST", "PUT", "PATCH", "DELETE"}

// WithMethod is a geeneric wrapper around a generic handler, forcing the a HTTP verb
func (l *Server[ConfigT]) WithMethod(method string, handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	// #StephenCurrying
	return func(w http.ResponseWriter, req *http.Request) {
		for _, m := range Methods {
			// a method matches
			if m == method {
				res, err := handler(w, req, l.Config)
				if err != nil {
					l.ResponseWriter(res, w)
				}
				slog.Error("Handler error", "error", err.Error())
				return
			}
		}
		// no method matched the one provided over the array of available methods
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method %s not allowed", req.Method)
	}
}

// Get is a wrapper around a generic handler, forcing the GET HTTP verb
func (l *Server[ConfigT]) Get(handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	return l.WithMethod("GET", handler)
}

// Post is a wrapper around a generic handler, forcing the POST HTTP verb
func (l *Server[ConfigT]) Post(handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	return l.WithMethod("POST", handler)
}

// Put is a wrapper around a generic handler, forcing the PUT HTTP verb
func (l *Server[ConfigT]) Put(handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	return l.WithMethod("PUT", handler)
}

// Patch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func (l *Server[ConfigT]) Patch(handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	return l.WithMethod("PATCH", handler)
}

// Delete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func (l *Server[ConfigT]) Delete(handler Handler[ConfigT]) func(http.ResponseWriter, *http.Request) {
	return l.WithMethod("DELETE", handler)
}
