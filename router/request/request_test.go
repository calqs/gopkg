package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Cabane123 struct {
	Cabane    int `query:"cabane"`
	NoTag     string
	PtrString *string `query:"ptr_string"`
}

func TestICanExtractQueryStringFromRequest(t *testing.T) {
	t.Run("simple query with simple struct", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodGet, "/test?cabane=123&notag=ok", nil)
		res, err := ExtractData[Cabane123](req1)
		assert.NoError(t, err)
		assert.Equal(t, Cabane123{Cabane: 123, NoTag: "ok", PtrString: nil}, *res)
	})
	t.Run("we can fill pointers", func(t *testing.T) {
		req2 := httptest.NewRequest(http.MethodGet, "/test?ptr_string=ok", nil)
		res, err := ExtractData[Cabane123](req2)
		assert.NoError(t, err)
		str := "ok"
		assert.Equal(t, Cabane123{Cabane: 0, NoTag: "", PtrString: &str}, *res)
	})
}
