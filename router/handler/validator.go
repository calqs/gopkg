package handler

import (
	"strings"

	"github.com/calqs/gopkg/router/response"
)

type DataObjectValidation interface {
	Validate() error
}

func WithValidation[TypeT DataObjectValidation](
	cb func(*Request[TypeT]) response.Response,
) func(*Request[TypeT]) response.Response {
	return func(req *Request[TypeT]) response.Response {
		if err := (*req.Params).Validate(); err != nil {
			return response.BadRequest(strings.ReplaceAll(err.Error(), "\n", ", "), err)
		}
		return cb(req)
	}
}
