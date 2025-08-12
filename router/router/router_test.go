package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/calqs/gopkg/router/response"
	"github.com/calqs/gopkg/router/testutil"
	"github.com/stretchr/testify/assert"
)

func TestICanDoRouting(t *testing.T) {
	type config struct{}
	router := &Router[config]{
		mux: http.NewServeMux(),
	}
	rw := testutil.NewStringResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router.HandleFuncs(
		"/test",
		HandleGet(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(`test_get`)
		}),
		HandlePost(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(`test_post`)
		}),
	)
	server := httptest.NewServer(router.mux)
	t.Cleanup(func() {
		server.Close()
	})
	t.Run("post & get routes: test post should be OK", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodPost, "/test", nil)
		rec1 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec1, req1)
		assert.Equal(t, http.StatusOK, rec1.Code)
		assert.Equal(t, strings.TrimSpace(rec1.Body.String()), "test_post")
	})
	t.Run("post & get routes: test get should be OK", func(t *testing.T) {
		req2 := httptest.NewRequest(http.MethodGet, "/test", nil)
		rec2 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec2, req2)
		assert.Equal(t, http.StatusOK, rec2.Code)
		assert.Equal(t, strings.TrimSpace(rec2.Body.String()), "test_get")
	})
	t.Run("post & get routes: test delete should be 405", func(t *testing.T) {
		req3 := httptest.NewRequest(http.MethodDelete, "/test", nil)
		rec3 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec3, req3)
		assert.Equal(t, http.StatusMethodNotAllowed, rec3.Code)
		httpErr := response.HttpError{}
		assert.NoError(t, json.Unmarshal(rec3.Body.Bytes(), &httpErr))
		assert.Equal(t, fmt.Sprintf(FormatMethodNotAllowed, DELETE.String(), "/test"), httpErr.Message)
	})
}

func TestAllMethod(t *testing.T) {
	type config struct {
		cabane int
	}
	type resp struct {
		Message string `json:"message"`
	}
	rw := testutil.NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router := NewRouter(&config{123})
	router.HandleFuncs(
		"/test/methods",
		HandleGet(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(&resp{"test_get"})
		}),
		HandlePost(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(&resp{"test_post"})
		}),
		HandlePut(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(&resp{"test_put"})
		}),
		HandlePatch(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(&resp{"test_patch"})
		}),
		HandleDelete(func(r *http.Request, ct *config) response.Response {
			return rw.WithAnyData(&resp{"test_delete"})
		}),
	)
	t.Run("get & post & put & patch & delete: get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test/methods", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"test_get"}, res)
	})
	t.Run("get & post & put & patch & delete: post", func(t *testing.T) {
		rw = rw.WithResponser(func(b []byte, w http.ResponseWriter) {
			w.WriteHeader(201)
			w.Write(b)
		})
		req := httptest.NewRequest(http.MethodPost, "/test/methods", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"test_post"}, res)
	})
	t.Run("get & post & put & patch & delete: put", func(t *testing.T) {
		rw = rw.WithResponser(func(b []byte, w http.ResponseWriter) {
			w.WriteHeader(http.StatusAccepted)
			w.Write(b)
		})
		req := httptest.NewRequest(http.MethodPut, "/test/methods", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusAccepted, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"test_put"}, res)
	})
	t.Run("get & post & put & patch & delete: patch", func(t *testing.T) {
		rw = rw.WithResponser(func(b []byte, w http.ResponseWriter) {
			w.WriteHeader(http.StatusNotModified)
			w.Write(b)
		})
		req := httptest.NewRequest(http.MethodPatch, "/test/methods", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotModified, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"test_patch"}, res)
	})
	t.Run("get & post & put & patch & delete: delete", func(t *testing.T) {
		rw = rw.WithResponser(func(b []byte, w http.ResponseWriter) {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		})
		req := httptest.NewRequest(http.MethodDelete, "/test/methods", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"test_delete"}, res)
	})
}
