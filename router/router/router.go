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
	tree        []*Router
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
		middlewares: middlewares.NewAPIMiddlewaresFromMux(),
		options:     serverOpts,
		tree:        make([]*Router, 0),
	}
}

func (swm *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prw := proxyResponseWriter{rw: w}
	for _, leaf := range swm.tree {
		leaf.ServeHTTP(&prw, r)
		if prw.eitherWriteWasCalled() {
			return
		}
	}
	swm.mux.ServeHTTP(w, r)
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
			swm.middlewares.MakeChain(mh).ServeHTTP(prw, req)
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
	for _, leaf := range swm.tree {
		leaf.Use(handlers...)
	}
}

func (swm *Router) Group(pattern string) *Router {
	opts := swm.options
	WithBaseURL(path.Join(opts.BaseURL, pattern))(&opts)
	leaf := &Router{
		mux:         swm.mux,
		ctx:         swm.ctx,
		middlewares: swm.middlewares.Clone(),
		options:     opts,
		tree:        make([]*Router, 0),
	}
	swm.tree = append(swm.tree, leaf)
	return leaf
}
