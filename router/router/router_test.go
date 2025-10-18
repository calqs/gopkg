package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/calqs/gopkg/router/handler"
	"github.com/calqs/gopkg/router/public"
	"github.com/calqs/gopkg/router/response"
	"github.com/stretchr/testify/assert"
)

type fake_req struct{}

func TestICanDoRouting(t *testing.T) {
	router := &Router{
		mux: http.NewServeMux(),
	}

	rw := NewStringResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router.Handle(
		"/test",
		public.Get(func(p *handler.Request[fake_req]) response.Response {
			return rw.WithAnyData("test_get")
		}),
		public.Post(func(p *handler.Request[fake_req]) response.Response {
			return rw.WithAnyData("test_post")
		}),
	)

	t.Run("post & get routes: test post should be OK", func(t *testing.T) {
		server := httptest.NewServer(router.mux)
		t.Cleanup(func() {
			server.Close()
		})
		req1 := httptest.NewRequest(http.MethodPost, "/test", nil)
		rec1 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec1, req1)
		assert.Equal(t, http.StatusOK, rec1.Code)
		assert.Equal(t, "test_post", strings.TrimSpace(rec1.Body.String()))
	})
	t.Run("post & get routes: test get should be OK", func(t *testing.T) {
		req2 := httptest.NewRequest(http.MethodGet, "/test", nil)
		rec2 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec2, req2)
		assert.Equal(t, http.StatusOK, rec2.Code)
		assert.Equal(t, "test_get", strings.TrimSpace(rec2.Body.String()))
	})
	t.Run("post & get routes: test delete should be 405", func(t *testing.T) {
		req3 := httptest.NewRequest(http.MethodDelete, "/test", nil)
		rec3 := httptest.NewRecorder()
		router.mux.ServeHTTP(rec3, req3)
		assert.Equal(t, http.StatusMethodNotAllowed, rec3.Code)
		httpErr := response.HTTPError{}
		assert.NoError(t, json.Unmarshal(rec3.Body.Bytes(), &httpErr))
		assert.Equal(t, fmt.Sprintf(FormatMethodNotAllowed, http.MethodDelete, "/test"), httpErr.Message)
	})
	t.Run("with options: base path", func(t *testing.T) {
		nrt := NewRouter(t.Context(), WithBaseURL("/cabane"))
		nrt.Handle(
			"/123",
			public.Get(func(p *handler.Request[fake_req]) response.Response {
				return rw.WithAnyData("test_get_123")
			}),
		)
		nrt.Handle(
			"/",
			public.Get(func(p *handler.Request[fake_req]) response.Response {
				return rw.WithAnyData("test_get_slash")
			}),
		)
		{
			req := httptest.NewRequest(http.MethodGet, "/cabane/123", nil)
			rec := httptest.NewRecorder()
			nrt.mux.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_get_123", strings.TrimSpace(rec.Body.String()))
		}
		{
			req := httptest.NewRequest(http.MethodGet, "/cabane", nil)
			rec := httptest.NewRecorder()
			nrt.mux.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_get_slash", strings.TrimSpace(rec.Body.String()))
		}
	})
}

func TestAllMethod(t *testing.T) {
	type resp struct {
		Message string `json:"message"`
	}
	rw := NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	router := NewRouter(t.Context())
	router.Handle(
		"/test/methods",
		public.Get(func(p *handler.Request[fake_req]) *FuncResponse {
			return rw.WithAnyData(&resp{"test_get"})
		}),
		public.Post(func(p *handler.Request[fake_req]) *FuncResponse {
			return rw.WithAnyData(&resp{"test_post"})
		}),
		public.Put(func(p *handler.Request[fake_req]) *FuncResponse {
			return rw.WithAnyData(&resp{"test_put"})
		}),
		public.Patch(func(p *handler.Request[fake_req]) *FuncResponse {
			return rw.WithAnyData(&resp{"test_patch"})
		}),
		public.Delete(func(p *handler.Request[fake_req]) *FuncResponse {
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
		rw := NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
			rw.WriteHeader(200)
			rw.Write(data)
		})
		type getReq struct {
			Cabane string `query:"cabane"`
		}
		type resp struct {
			Message string `json:"message"`
		}
		router := NewRouter(t.Context())
		router.Handle(
			"/test/request",
			public.Get(func(p *handler.Request[getReq]) *FuncResponse {
				return rw.WithAnyData(&resp{p.Params.Cabane})
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
		rw := NewJsonResponse(func(data []byte, rw http.ResponseWriter) {
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
		router := NewRouter(t.Context())
		router.Handle(
			"/test/request",
			public.Post(func(p *handler.Request[request]) *FuncResponse {
				cbn, err := strconv.Atoi(p.Params.Cabane)
				assert.NoError(t, err)
				return rw.WithAnyData(&resp{cbn, p.Params.Dog, p.Params.Amount})
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

func Test_I_Can_Route_Groups(t *testing.T) {
	fakeMiddleware := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.Header.Add("japan", "ikitai")
			h.ServeHTTP(w, req)
		})
	}
	br := NewRouter(t.Context(), WithBaseURL("/test"))
	rw := NewStringResponse(func(data []byte, rw http.ResponseWriter) {
		rw.WriteHeader(200)
		rw.Write(data)
	})
	br.Handle(
		"/",
		public.Get(func(p *handler.Request[fake_req]) *FuncResponse {
			assert.Empty(t, p.Request.Header.Get("japan"))
			return rw.WithAnyData("test_get")
		}),
	)
	t.Run("very simple get test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		rec := httptest.NewRecorder()
		br.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "test_get", rec.Body.String())
	})
	t.Run("with nested group", func(t *testing.T) {
		gr := br.Group("/cabane")
		gr.Handle(
			"/",
			public.Get(func(p *handler.Request[fake_req]) *FuncResponse {
				if p.Request.Header.Get("japan") != "" {
					t.Fail()
				}
				return rw.WithAnyData("test_nested_cabane")
			}),
		)
		{
			req := httptest.NewRequest(http.MethodGet, "/test/cabane", nil)
			rec := httptest.NewRecorder()
			gr.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_nested_cabane", rec.Body.String())
		}
		gr.Handle(
			"/123",
			public.Get(func(p *handler.Request[fake_req]) *FuncResponse {
				return rw.WithAnyData("test_nested_cabane_123")
			}),
		)
		{
			req := httptest.NewRequest(http.MethodGet, "/test/cabane/123", nil)
			rec := httptest.NewRecorder()
			gr.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_nested_cabane_123", rec.Body.String())
		}
		{
			// check base route is still good
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()
			br.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_get", rec.Body.String())
		}
	})
	t.Run("with nested group and middleware", func(t *testing.T) {
		{
			gr := br.Group("/future")
			gr.Use(fakeMiddleware)
			gr.Handle("/", public.Get(func(p *handler.Request[fake_req]) *FuncResponse {
				return rw.WithAnyData(p.Request.Header.Get("japan"))
			}))
			req := httptest.NewRequest(http.MethodGet, "/test/future", nil)
			rec := httptest.NewRecorder()
			gr.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "ikitai", rec.Body.String())
		}
		{
			// check base route is still good
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()
			br.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test_get", rec.Body.String())
		}
	})
}
