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
	handlers    map[string][]http.Handler
	served      bool
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
		handlers:    make(map[string][]http.Handler, 0),
		served:      false,
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
	if _, ok := swm.handlers[pattern]; !ok {
		swm.handlers[pattern] = make([]http.Handler, 0)
	}
	swm.handlers[pattern] = append(swm.handlers[pattern], mhs...)
}

func (swm *Router) Use(handlers ...func(http.Handler) http.Handler) {
	swm.middlewares.Use(handlers...)
	for _, leaf := range swm.tree {
		leaf.Use(handlers...)
	}
}

func (swm *Router) Load() *Router {
	for _, leaf := range swm.tree {
		leaf.Load()
	}
	for pattern, handlers := range swm.handlers {
		swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
			prw := &proxyResponseWriter{w, false, false}
			if !areReqResOk(prw, req) {
				return
			}
			for _, mh := range handlers {
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
	return swm
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
		handlers:    make(map[string][]http.Handler, 0),
		served:      false,
	}
	swm.tree = append(swm.tree, leaf)
	return leaf
}
