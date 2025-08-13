package response

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type HttpError struct {
	Code    int
	Err     error
	Message string
	rw      http.ResponseWriter
}

func (e *HttpError) Error() string {
	return e.Err.Error()
}

func (e *HttpError) Send(w http.ResponseWriter) {
	slog.Error("API error", "error", e.Error(), "code", e.Code, "message", e.Message)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"code": %d, "message": "%s"}`, e.Code, e.Message)
}

func NewError(code int, msg string, errs ...error) *HttpError {
	if len(errs) == 0 {
		errs = []error{errors.New(msg)}
	}
	return &HttpError{
		Code:    code,
		Message: msg,
		Err:     errs[0],
	}
}

func BadRequest(msg string, errs ...error) *HttpError {
	return NewError(http.StatusBadRequest, msg, errs...)
}

func Unauthorized(msg string, errs ...error) *HttpError {
	return NewError(http.StatusUnauthorized, msg, errs...)
}

func UnprocessableEntity(msg string, errs ...error) *HttpError {
	return NewError(http.StatusUnprocessableEntity, msg, errs...)
}

func InternalServerError(msg string, errs ...error) *HttpError {
	return NewError(http.StatusInternalServerError, msg, errs...)
}

func MethodNotAllowed(msg string, errs ...error) *HttpError {
	return NewError(http.StatusMethodNotAllowed, msg, errs...)
}
