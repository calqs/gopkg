package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeHandler struct {
	serveHTTP func(w http.ResponseWriter, r *http.Request)
}

func (fk fakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fk.serveHTTP(w, r)
}

func Test_Add_Middlewares(t *testing.T) {
	type cabane string
	req1 := httptest.NewRequest(http.MethodPost, "/test", nil)
	rec1 := httptest.NewRecorder()
	count := 0

	mids := NewAPIMiddlewaresFromMux(fakeHandler{
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, 456, r.Context().Value(cabane("cabane")))
			w.WriteHeader(999)
		},
	})

	goesZero := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 0, count)
		count++
	}

	goesFirst := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 1, count)
		count++
	}

	goesSecond := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), cabane("cabane"), 123))
			assert.Equal(t, 2, count)
			count++
			h.ServeHTTP(w, r)
		})
	}

	goesThird := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, 123, r.Context().Value(cabane("cabane")))
			r = r.WithContext(context.WithValue(r.Context(), cabane("cabane"), 456))
			assert.Equal(t, 3, count)
			count++
			h.ServeHTTP(w, r)
			assert.Equal(t, 999, w.(*httptest.ResponseRecorder).Code)
		})
	}

	goesFourth := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 456, r.Context().Value(cabane("cabane")))
		assert.Equal(t, 999, w.(*httptest.ResponseRecorder).Code)
		assert.Equal(t, 4, count)
		count++
	}

	goesFifth := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 5, count)
		count++
	}

	mids.UseAfter(goesFourth)
	mids.UseAfter(goesFifth)
	mids.Use(goesSecond, goesThird)
	mids.UseBefore(goesFirst)
	mids.UseBefore(goesZero)
	mids.ServeHTTP(rec1, req1)
}
