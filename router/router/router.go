package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/calqs/gopkg/router/handler"
	"github.com/calqs/gopkg/router/middlewares"
	"github.com/calqs/gopkg/router/response"
)

type Router struct {
	mux         *http.ServeMux
	middlewares *middlewares.APIMiddlewares
	ctx         context.Context
	options     Options
}

func areReqResOk(w http.ResponseWriter, req *http.Request) bool {
	if w == nil {
		panic(ErrNilRequestOrResponse.Error())
	}
	if req == nil {
		response.
			InternalServerError(ErrNilRequestOrResponse.Error(), fmt.Errorf("routeIt: *http.Request or http.ResponseWriter: %w", ErrNilRequestOrResponse)).
			Send(w)
		return false
	}
	return true
}

func (swm *Router) routeIt(w http.ResponseWriter, req *http.Request, mh handler.MethodHandler) {
	if !areReqResOk(w, req) {
		return
	}
	res := mh.Handler(req)
	res.Send(w)
}

func (swm *Router) Handle(pattern string, mhs ...handler.MethodHandler) {
	// have to clean path, at least for security reasons
	pattern = CleanPath(swm.options.BaseURL + CleanPath(pattern))
	swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if !areReqResOk(w, req) {
			return
		}
		for _, mh := range mhs {
			if req.Method != mh.Method.String() {
				continue
			}
			if pattern != req.RequestURI {
				continue
			}
			swm.routeIt(w, req, mh)
			return
		}
		response.
			MethodNotAllowed(fmt.Sprintf(FormatMethodNotAllowed, req.Method, pattern)).
			Send(w)
	})
}

func (swm *Router) Use(handlers ...func(http.Handler) http.Handler) {
	swm.middlewares.Use(handlers...)
}

func (swm *Router) GetHttpHandler() http.Handler {
	return swm.middlewares
}

func (swm *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	swm.middlewares.ServeHTTP(w, r)
}

func NewRouter(ctx context.Context, opts ...OptionFunc) *Router {
	mux := http.NewServeMux()
	serverOpts := Options{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&serverOpts)
	}
	return &Router{
		mux:         mux,
		ctx:         ctx,
		middlewares: middlewares.NewAPIMiddlewaresFromMux(mux),
		options:     serverOpts,
	}
}
