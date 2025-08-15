package middlewares

import (
	"net/http"
)

type APIMiddlewares struct {
	chain http.Handler
}

func NewAPIMiddlewaresFromMux(handler http.Handler) *APIMiddlewares {
	return &APIMiddlewares{handler}
}

func (a *APIMiddlewares) Use(handlers ...func(http.Handler) http.Handler) {
	for i := range len(handlers) {
		a.chain = handlers[len(handlers)-1-i](a.chain)
	}
}

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

func (a *APIMiddlewares) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.chain.ServeHTTP(w, r)
}
