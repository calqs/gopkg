package response

import (
	"net/http"
)

type Response interface {
	Send(http.ResponseWriter)
}

// HTTPResponse is a standard HTTP response, meant for 2xx, for example
type HTTPResponse struct {
	Data       []byte
	StatusCode int
}

func (hr *HTTPResponse) Send(w http.ResponseWriter) {
	w.WriteHeader(hr.StatusCode)
	w.Write(hr.Data)
}

// RedirectResponse is a "Location" header driven HTTP response, meant for 3xx
type RedirectResponse struct {
	location   string
	StatusCode int
}

func (rr *RedirectResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Location", rr.location)
	w.WriteHeader(rr.StatusCode)
}
