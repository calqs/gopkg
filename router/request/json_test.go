package request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fake_request struct {
	Cabane int `json:"cabane"`
}

func Test_JsonBodyRequest(t *testing.T) {
	t.Run("with nil req", func(t *testing.T) {
		req, err := JsonBodyRequest[fake_request](nil)
		assert.Nil(t, err)
		assert.Nil(t, req)
	})
	t.Run("with a body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test/methods", strings.NewReader(`{"cabane": 123}`))
		req.Header.Set("content-type", "application/json")
		res, err := JsonBodyRequest[fake_request](req)
		assert.NoError(t, err)
		assert.Equal(t, fake_request{123}, *res)
	})
	t.Run("with a body and a bad json field type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test/methods", strings.NewReader(`{"cabane": "string_123"}`))
		req.Header.Set("content-type", "application/json")
		res, err := JsonBodyRequest[fake_request](req)
		assert.ErrorIs(t, err, ErrPayloadWrongFieldType)
		assert.ErrorIs(t, err, ErrPayloadMalformed)
		assert.Nil(t, res)
	})
	t.Run("with a body and a bad json shape", func(t *testing.T) {
		// missing trailing }
		req := httptest.NewRequest(http.MethodGet, "/test/methods", strings.NewReader(`{"cabane": 123, "dog": "suzie"`))
		req.Header.Set("content-type", "application/json")
		res, err := JsonBodyRequest[fake_request](req)
		assert.ErrorIs(t, err, ErrPayloadWrongShape)
		assert.ErrorIs(t, err, ErrPayloadMalformed)
		assert.Nil(t, res)
	})
}
