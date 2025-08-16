package response

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// HTTPError implements Response interface, for errors only
type HTTPError struct {
	Code    int
	Err     error
	Message string
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func (e *HTTPError) Send(w http.ResponseWriter) {
	slog.Error("API error", "error", e.Error(), "code", e.Code, "message", e.Message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	fmt.Fprintf(w, `{"code": %d, "message": "%s"}`, e.Code, e.Message)
}

func NewError(code int, msg string, errs ...error) *HTTPError {
	if len(errs) == 0 {
		errs = []error{errors.New(msg)}
	}
	return &HTTPError{
		Code:    code,
		Message: msg,
		Err:     errs[0],
	}
}
