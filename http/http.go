package http

import (
	"github.com/calqs/http/response"
	"github.com/calqs/http/server"
)

func NewJSONServer[ConfigT any]() *server.Server[ConfigT] {
	return &server.Server[ConfigT]{
		ResponseWriter: response.JsonResponseWriter,
	}
}
