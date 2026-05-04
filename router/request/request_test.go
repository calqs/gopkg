package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	t.Run("we can fill time.Time", func(t *testing.T) {
		req2 := httptest.NewRequest(http.MethodGet, "/test?created_at=2026-01-01T00:00:00Z", nil)
		res, err := ExtractData[struct {
			CreatedAt time.Time `query:"created_at"`
		}](req2)
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), res.CreatedAt)
	})
}
