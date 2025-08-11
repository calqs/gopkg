package router

import "net/http"

type withWriteRecorder struct {
	http.ResponseWriter
	wrote bool
}

func (tw *withWriteRecorder) WriteHeader(code int) {
	tw.wrote = true
	tw.ResponseWriter.WriteHeader(code)
}

func (tw *withWriteRecorder) Write(b []byte) (int, error) {
	tw.wrote = true
	return tw.ResponseWriter.Write(b)
}

func (tw *withWriteRecorder) Wrote() bool {
	return tw.wrote
}
