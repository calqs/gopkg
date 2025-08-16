package middlewares

import (
	"net/http"
)

// APIMiddlewares represents a middleware system for a classic API structure (eg: json API etc...).
// Those middlewares are nothing more than http.Handlers called before or after the actual
// endpoint http.Handler
type APIMiddlewares struct {
	chain http.Handler
}

// NewAPIMiddlewaresFromMux creates a new APIMiddlewares from a serverMux
func NewAPIMiddlewaresFromMux(serverMux http.Handler) *APIMiddlewares {
	return &APIMiddlewares{serverMux}
}

// Use adds one or more middleware to the pile of http.Handlers acting as middlewares.
// Those middlewares will receive an http.Handler as parameter, allowing them to
// modifying some parts of the http.Request, add stuff to the context, log before and after the endpoint handler.
// Those are some of the many possile uses of middlewares.
func (a *APIMiddlewares) Use(handlers ...func(http.Handler) http.Handler) {
	for i := range len(handlers) {
		a.chain = handlers[len(handlers)-1-i](a.chain)
	}
}

// UseBefore adds one or more handlers to the pile of middlewares, forcing them to be called
// BEFORE the actual endpoint handler.
// UseBefore works as a LIFO pile, last declared middleware will be the first to be used.
func (a *APIMiddlewares) UseBefore(handlers ...func(http.ResponseWriter, *http.Request)) {
	for _, handler := range handlers {
		a.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handler(w, r)
				h.ServeHTTP(w, r)
			})
		})
	}
}

// UseAfter adds one or more handlers to the pile of middlewares, forcing them to be called
// AFTER the actual endpoint handler.
// UseAfter works as a FIFO pile, first declared middleware will be the first to be used.
func (a *APIMiddlewares) UseAfter(handlers ...func(http.ResponseWriter, *http.Request)) {
	for _, handler := range handlers {
		a.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.ServeHTTP(w, r)
				handler(w, r)
			})
		})
	}
}

// ServeHTTP implements http.Handler
func (a *APIMiddlewares) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.chain.ServeHTTP(w, r)
}
