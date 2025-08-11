package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestICanDoRouting(t *testing.T) {
	type config struct{}
	router := &Router[config]{
		baseRouter: &baseRouter[config]{
			Config: &config{},
			ResponseWriter: func(res any, rw http.ResponseWriter) {
				rw.WriteHeader(200)
				rw.Write([]byte("ok"))
			},
		},
		mux: http.NewServeMux(),
	}
	// mux.HandleFunc("/test", router.HandleGet(func(w http.ResponseWriter, r *http.Request, c ...*config) (any, *response.HttpError) {
	// 	return "test", nil
	// }))
	// mux.HandleFunc("/test", router.HandlePost(func(w http.ResponseWriter, r *http.Request, c ...*config) (any, *response.HttpError) {
	// 	return "test", nil
	// }))
	server := httptest.NewServer(router.mux)
	t.Cleanup(func() {
		server.Close()
	})
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	router.mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d; want %d", rec.Code, http.StatusOK)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "ok" {
		t.Fatalf("body = %q; want %q", got, "ok")
	}
}
