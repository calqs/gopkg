package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Cabane123 struct {
	Cabane int `query:"cabane"`
	NoTag  string
}

func TestICanExtractQueryStringFromRequest(t *testing.T) {
	t.Run("simple query with simple struct", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodGet, "/test?cabane=123&notag=ok", nil)
		res, err := ExtractData[Cabane123](req1)
		assert.NoError(t, err)
		assert.Equal(t, Cabane123{Cabane: 123, NoTag: "ok"}, *res)
	})
}
