package router

import (
	"github.com/calqs/gopkg/router/router"
)

func NewJsonRouter[ConfigT any]() *router.Router[ConfigT] {
	return router.NewRouter[ConfigT](JsonResponseWriter)
}
