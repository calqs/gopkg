package response

import "net/http"

type Response interface {
	Send(http.ResponseWriter)
}

type HTTPResponse struct {
	data       []byte
	statusCode int
}

func (hr *HTTPResponse) Send(w http.ResponseWriter) {
	w.WriteHeader(hr.statusCode)
	w.Write(hr.data)
}
