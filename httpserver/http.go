package httpserver

import (
	"github.com/calqs/gopkg/httpserver/response"
	"github.com/calqs/gopkg/httpserver/server"
)

func NewJSONServer[ConfigT any]() *server.Server[ConfigT] {
	return &server.Server[ConfigT]{
		ResponseWriter: response.JsonResponseWriter,
	}
}
