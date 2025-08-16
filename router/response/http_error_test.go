package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPError_Error(t *testing.T) {
	err := errors.New("something went wrong")
	httpErr := &HTTPError{Err: err}
	if got := httpErr.Error(); got != "something went wrong" {
		t.Errorf("Error() = %q, want %q", got, err.Error())
	}
}

func TestHTTPError_Send(t *testing.T) {
	rr := httptest.NewRecorder()
	httpErr := &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "bad input",
		Err:     errors.New("bad input"),
	}

	httpErr.Send(rr)

	// Check Content-Type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}

	// Parse response body
	var resp map[string]any
	if err := json.NewDecoder(bytes.NewReader(rr.Body.Bytes())).Decode(&resp); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}

	if int(resp["code"].(float64)) != http.StatusBadRequest {
		t.Errorf("code = %v, want %d", resp["code"], http.StatusBadRequest)
	}
	if resp["message"] != "bad input" {
		t.Errorf("message = %q, want %q", resp["message"], "bad input")
	}
}

func TestNewError_DefaultError(t *testing.T) {
	errObj := NewError(400, "oops")
	if errObj.Err.Error() != "oops" {
		t.Errorf("default Err = %q, want %q", errObj.Err.Error(), "oops")
	}
	if errObj.Code != 400 {
		t.Errorf("Code = %d, want 400", errObj.Code)
	}
	if errObj.Message != "oops" {
		t.Errorf("Message = %q, want oops", errObj.Message)
	}
}

func TestNewError_WithCustomError(t *testing.T) {
	customErr := errors.New("custom")
	errObj := NewError(500, "ignored", customErr)
	if errObj.Err != customErr {
		t.Errorf("Err = %v, want %v", errObj.Err, customErr)
	}
	if errObj.Message != "ignored" {
		t.Errorf("Message = %q, want ignored", errObj.Message)
	}
}

func TestHelpers(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(string, ...error) *HTTPError
		code    int
		message string
		custom  error
	}{
		{"BadRequest", BadRequest, http.StatusBadRequest, "bad", nil},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized, "unauth", nil},
		{"UnprocessableEntity", UnprocessableEntity, http.StatusUnprocessableEntity, "invalid", nil},
		{"InternalServerError", InternalServerError, http.StatusInternalServerError, "fail", nil},
		{"MethodNotAllowed", MethodNotAllowed, http.StatusMethodNotAllowed, "nope", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errObj := tt.fn(tt.message)
			if errObj.Code != tt.code {
				t.Errorf("Code = %d, want %d", errObj.Code, tt.code)
			}
			if errObj.Message != tt.message {
				t.Errorf("Message = %q, want %q", errObj.Message, tt.message)
			}
			if errObj.Err.Error() != tt.message {
				t.Errorf("Err.Error() = %q, want %q", errObj.Err.Error(), tt.message)
			}
		})
	}
}
