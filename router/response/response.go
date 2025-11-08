package response

import (
	"net/http"
)

type Response interface {
	Send(http.ResponseWriter)
	SetHeader(string, string)
}

type ResponseHeaders struct {
	Headers [][2]string
}

func (rh *ResponseHeaders) SetHeader(key, value string) {
	rh.Headers = append(rh.Headers, [2]string{key, value})
}

func (rh *ResponseHeaders) SetCookie(cookie http.Cookie) {
	if v := cookie.String(); v != "" {
		rh.SetHeader("Set-Cookie", cookie.String())
	}
}

func (rh *ResponseHeaders) WriteHeaders(w http.ResponseWriter) {
	for _, header := range rh.Headers {
		w.Header().Set(header[0], header[1])
	}
}

// HTTPResponse is a standard HTTP response, meant for 2xx, for example
type HTTPResponse struct {
	ResponseHeaders
	Data       []byte
	StatusCode int
}

func (hr *HTTPResponse) Send(w http.ResponseWriter) {
	hr.WriteHeaders(w)
	w.WriteHeader(hr.StatusCode)
	w.Write(hr.Data)
}

// RedirectResponse is a "Location" header driven HTTP response, meant for 3xx
type RedirectResponse struct {
	ResponseHeaders
	location   string
	StatusCode int
}

func (rr *RedirectResponse) Send(w http.ResponseWriter) {
	rr.WriteHeaders(w)
	w.Header().Set("Location", rr.location)
	w.WriteHeader(rr.StatusCode)
}
