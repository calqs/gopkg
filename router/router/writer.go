package router

import (
	"net/http"
)

type proxyResponseWriter struct {
	rw                http.ResponseWriter
	writeHeaderCalled bool
	writeCalled       bool
}

func (prw proxyResponseWriter) Header() http.Header {
	return prw.rw.Header()
}

func (prw *proxyResponseWriter) Write(b []byte) (int, error) {
	prw.writeCalled = true
	return prw.rw.Write(b)
}

func (prw *proxyResponseWriter) WriteHeader(statusCode int) {
	prw.writeHeaderCalled = true
	prw.rw.WriteHeader(statusCode)
}

func (prw proxyResponseWriter) eitherWriteWasCalled() bool {
	return prw.writeCalled || prw.writeHeaderCalled
}
