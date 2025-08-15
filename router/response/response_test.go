package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPResponse_Send(t *testing.T) {
	rec := httptest.NewRecorder()
	resp := &HTTPResponse{
		data:       []byte("hello"),
		statusCode: http.StatusCreated,
	}
	resp.Send(rec)
	assert.Equal(t, http.StatusCreated, rec.Code, "status code should match")
	assert.Equal(t, "hello", rec.Body.String(), "body should match")
}

type testTransformer struct{}

func (testTransformer) Transform(data any) []byte {
	return []byte("json:" + string(data.([]byte)))
}

func TestStatusOk_UsesTransformer(t *testing.T) {
	resp := StatusOk[testTransformer]([]byte("input"))
	assert.Equal(t, http.StatusOK, resp.statusCode, "status code should be HTTP 200 OK")
	assert.Equal(t, "json:input", string(resp.data), "data should be transformed correctly")
}
