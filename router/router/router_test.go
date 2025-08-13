package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/calqs/gopkg/router/response"
	"github.com/calqs/gopkg/router/testutil"
	"github.com/stretchr/testify/assert"
)

func TestICanDoRouting(t *testing.T) {
	router := &Router{
		mux: http.NewServeMux(),
	}
	type fake_req struct{}
	rw := testutil.NewStringResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router.Handle(
		"/test",
		Get(func(_ *fake_req, r *http.Request) response.Response {
			return rw.WithAnyData("test_get")
		}),
		Post(func(_ *fake_req, r *http.Request) response.Response {
			return rw.WithAnyData("test_post")
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
	type fake_req struct{}
	type resp struct {
		Message string `json:"message"`
	}
	rw := testutil.NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router := NewRouter()
	router.Handle(
		"/test/methods",
		Get(func(d *fake_req, r *http.Request) *testutil.FuncResponse {
			return rw.WithAnyData(&resp{"test_get"})
		}),
		Post(func(d *fake_req, r *http.Request) *testutil.FuncResponse {
			return rw.WithAnyData(&resp{"test_post"})
		}),
		Put(func(d *fake_req, r *http.Request) *testutil.FuncResponse {
			return rw.WithAnyData(&resp{"test_put"})
		}),
		Patch(func(d *fake_req, r *http.Request) *testutil.FuncResponse {
			return rw.WithAnyData(&resp{"test_patch"})
		}),
		Delete(func(d *fake_req, r *http.Request) *testutil.FuncResponse {
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

func TestComplexRoutesWithParams(t *testing.T) {
	t.Run("get request with query params", func(t *testing.T) {
		rw := testutil.NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
			rw.WriteHeader(200)
			rw.Write(data)
		})
		type getReq struct {
			Cabane string `query:"cabane"`
		}
		type resp struct {
			Message string `json:"message"`
		}
		router := NewRouter()
		router.Handle(
			"/test/request",
			Get(func(d *getReq, r *http.Request) *testutil.FuncResponse {
				return rw.WithAnyData(&resp{d.Cabane})
			}),
		)
		req := httptest.NewRequest(http.MethodGet, "/test/request?cabane=123", nil)
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{"123"}, res)
	})

	t.Run("post request with query params", func(t *testing.T) {
		rw := testutil.NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
			rw.WriteHeader(http.StatusCreated)
			rw.Write(data)
		})
		type request struct {
			Cabane string `query:"cabane"`
			Dog    string
			Amount int
		}
		type resp struct {
			Cabane int    `json:"cabane"`
			Dog    string `json:"dog"`
			Amount int    `json:"amount"`
		}
		router := NewRouter()
		router.Handle(
			"/test/request",
			Post(func(d *request, r *http.Request) *testutil.FuncResponse {
				cbn, err := strconv.Atoi(d.Cabane)
				assert.NoError(t, err)
				return rw.WithAnyData(&resp{cbn, d.Dog, d.Amount})
			}),
		)
		req := httptest.NewRequest(http.MethodPost, "/test/request?cabane=123", strings.NewReader(`{"dog": "suzie", "amount": 1}`))
		rec := httptest.NewRecorder()
		router.mux.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
		res := resp{}
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, resp{123, "suzie", 1}, res)
	})
}
