package router

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/calqs/gopkg/router/middlewares"
	"github.com/calqs/gopkg/router/response"
)

type Router struct {
	mux         *http.ServeMux
	middlewares *middlewares.APIMiddlewares
	ctx         context.Context
	options     Options
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

func (swm *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	swm.middlewares.ServeHTTP(w, r)
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

func (swm *Router) Handle(pattern string, mhs ...http.Handler) {
	// have to clean path, at least for security reasons
	pattern = CleanPath(swm.options.BaseURL + CleanPath(pattern))
	swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		prw := &proxyResponseWriter{w, false, false}
		if !areReqResOk(prw, req) {
			return
		}
		for _, mh := range mhs {
			if pattern != req.URL.Path {
				continue
			}
			mh.ServeHTTP(prw, req)
			if !prw.eitherWriteWasCalled() {
				continue
			}
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

func (swm *Router) Group(pattern string) *Router {
	opts := swm.options
	WithBaseURL(path.Join(opts.BaseURL, pattern))(&opts)
	return &Router{
		mux:         swm.mux,
		ctx:         swm.ctx,
		middlewares: middlewares.NewAPIMiddlewaresFromMux(swm.mux),
		options:     opts,
	}
}
