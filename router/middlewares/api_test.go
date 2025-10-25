package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeHandler struct {
	serveHTTP func(w http.ResponseWriter, r *http.Request)
}

type testCtx struct {
	cabane int
}

func (fk fakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fk.serveHTTP(w, r)
}

func Test_Add_Middlewares(t *testing.T) {
	type cabane string
	req1 := httptest.NewRequest(http.MethodPost, "/test", nil)
	rec1 := httptest.NewRecorder()
	count := 0
	mids := NewAPIMiddlewaresFromMux()

	goesZero := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("0 (USE BEFORE)", count)
		assert.Equal(t, 0, count)
		count++
	}

	goesFirst := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1 (USE BEFORE)", count)
		assert.Equal(t, 1, count)
		count++
	}

	goesSecond := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Context().Value(cabane("cabane")).(*testCtx).cabane = 123
			assert.Equal(t, 2, count)
			fmt.Println("2 (USE) before", count)
			count++
			h.ServeHTTP(w, r)
			// third mw goes right after this one, so count gets incremented once more
			assert.Equal(t, 4, count)
			fmt.Println("2 (USE) after", count)
		})
	}

	goesThird := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, &testCtx{123}, r.Context().Value(cabane("cabane")))
			r.Context().Value(cabane("cabane")).(*testCtx).cabane = 456
			assert.Equal(t, 3, count)
			fmt.Println("3 (USE) before", count)
			count++
			h.ServeHTTP(w, r)
			assert.Equal(t, 4, count)
			fmt.Println("3 (USE) after", count)
			assert.Equal(t, 999, w.(*httptest.ResponseRecorder).Code)
		})
	}

	goesFourth := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("4 (USE AFTER)", count)
		assert.Equal(t, &testCtx{456}, r.Context().Value(cabane("cabane")))
		assert.Equal(t, 999, w.(*httptest.ResponseRecorder).Code)
		assert.Equal(t, 4, count)
		count++
	}

	goesFifth := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("5 (USE AFTER)", count)
		assert.Equal(t, 5, count)
		count++
	}

	mids.UseAfter(goesFourth)
	mids.UseAfter(goesFifth)
	mids.Use(goesSecond, goesThird)
	mids.UseBefore(goesFirst)
	mids.UseBefore(goesZero)
	data := &testCtx{0}
	mids.MakeChain(fakeHandler{
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, &testCtx{456}, r.Context().Value(cabane("cabane")))
			w.WriteHeader(999)
		},
	}).ServeHTTP(rec1, req1.WithContext(context.WithValue(req1.Context(), cabane("cabane"), data)))
}
