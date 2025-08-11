package router

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/calqs/gopkg/router/middlewares"
)

type Router[ConfigT any] struct {
	*baseRouter[ConfigT]
	mux *http.ServeMux
	api *middlewares.API
}

func (swm *Router[ConfigT]) HandleFunc(pattern string, mh MethodHandler[ConfigT]) {
	swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == mh.Method.String() {
			res, err := mh.Handler(w, req, swm.Config)
			if err != nil {
				slog.Error("Handler error", "error", err.Error())
				err.Write(w)
				return
			}
			swm.ResponseWriter(res, w)
			return
		}
		w.WriteHeader(405)
		fmt.Fprintf(w, "could not find any route matching %s %s", req.Method, pattern)
	})
}

func (swm *Router[ConfigT]) HandleFuncs(pattern string, mhs ...MethodHandler[ConfigT]) {
	swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		for _, mh := range mhs {
			if req.Method != mh.Method.String() {
				continue
			}
			res, err := mh.Handler(w, req, swm.Config)
			if err != nil {
				slog.Error("Handler error", "error", err.Error())
				err.Write(w)
				return
			}
			swm.ResponseWriter(res, w)
			return
		}
		w.WriteHeader(405)
		fmt.Fprintf(w, "could not find any route matching %s %s", req.Method, pattern)
	})
}

// Get is a wrapper around a generic handler, forcing the GET HTTP verb
func (swm *Router[ConfigT]) Get(pattern string, handler Handler[ConfigT]) {
	swm.HandleFunc(pattern, HandleGet(handler))
}

// Post is a wrapper around a generic handler, forcing the POST HTTP verb
func (swm *Router[ConfigT]) Post(pattern string, handler Handler[ConfigT]) {
	swm.HandleFunc(pattern, HandlePost(handler))
}

// Put is a wrapper around a generic handler, forcing the PUT HTTP verb
func (swm *Router[ConfigT]) Put(pattern string, handler Handler[ConfigT]) {
	swm.HandleFunc(pattern, HandlePut(handler))
}

// Patch is a wrapper around a generic handler, forcing the PATCH HTTP verb
func (swm *Router[ConfigT]) Patch(pattern string, handler Handler[ConfigT]) {
	swm.HandleFunc(pattern, HandlePatch(handler))
}

// Delete is a wrapper around a generic handler, forcing the DELETE HTTP verb
func (swm *Router[ConfigT]) Delete(pattern string, handler Handler[ConfigT]) {
	swm.HandleFunc(pattern, HandleDelete(handler))
}

func (swm *Router[ConfigT]) Use(handlers ...func(http.Handler) http.Handler) {
	if swm.api == nil {
		swm.api = middlewares.NewAPIFromMux(swm.mux)
	}
	swm.api.Use(handlers...)
}

func (swm *Router[ConfigT]) GetHttpHandler() http.Handler {
	return swm.api
}

func NewRouter[ConfigT any](rw func(any, http.ResponseWriter)) *Router[ConfigT] {
	return &Router[ConfigT]{
		baseRouter: &baseRouter[ConfigT]{
			ResponseWriter: rw,
		},
		mux: http.NewServeMux(),
	}
}
