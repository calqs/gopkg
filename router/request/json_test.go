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
		_, err := JsonBodyRequest[fake_request](nil)
		assert.ErrorIs(t, err, ErrNilPointer)
		assert.ErrorContains(t, err, "http.Request")
	})
	t.Run("with a body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test/methods", strings.NewReader(`{"cabane": 123}`))
		res, err := JsonBodyRequest[fake_request](req)
		assert.NoError(t, err)
		assert.Equal(t, fake_request{123}, *res)
	})
}
