package router

import (
	"fmt"
	"net/http"

	"github.com/calqs/gopkg/router/middlewares"
	"github.com/calqs/gopkg/router/response"
)

type Router struct {
	mux *http.ServeMux
	api *middlewares.API
}

func (swm *Router) routeIt(w http.ResponseWriter, req *http.Request, mh MethodHandler) {
	res := mh.Handler(req)
	res.Send(w)
}

func (swm *Router) Handle(pattern string, mhs ...MethodHandler) {
	swm.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		for _, mh := range mhs {
			if req.Method != mh.Method.String() {
				continue
			}
			swm.routeIt(w, req, mh)
			return
		}
		w.WriteHeader(405)
		response.
			MethodNotAllowed(fmt.Sprintf(FormatMethodNotAllowed, req.Method, pattern)).
			Send(w)
	})
}

func (swm *Router) Use(handlers ...func(http.Handler) http.Handler) {
	if swm.api == nil {
		swm.api = middlewares.NewAPIFromMux(swm.mux)
	}
	swm.api.Use(handlers...)
}

func (swm *Router) GetHttpHandler() http.Handler {
	return swm.api
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}
