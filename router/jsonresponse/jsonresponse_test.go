package jsonresponse

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusOk(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusOk(map[string]string{"hello": "world"})
	resp.Send(rr)

	assert.Equal(t, http.StatusOK, rr.Code)

	var out map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, "world", out["hello"])
}

func TestStatusCreated(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusCreated(map[string]int{"id": 123})
	resp.Send(rr)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var out map[string]int
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, 123, out["id"])
}

func TestStatusAccepted(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusAccepted("accepted")
	resp.Send(rr)

	assert.Equal(t, http.StatusAccepted, rr.Code)

	var out string
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, "accepted", out)
}

func TestStatusNonAuthoritativeInfo(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusNonAuthoritativeInfo(true)
	resp.Send(rr)

	assert.Equal(t, http.StatusNonAuthoritativeInfo, rr.Code)

	var out bool
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.True(t, out)
}

func TestStatusNoContent(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusNoContent()
	resp.Send(rr)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
}

func TestStatusResetContent(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusResetContent()
	resp.Send(rr)

	assert.Equal(t, http.StatusResetContent, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
}

func TestStatusPartialContent(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusPartialContent([]int{1, 2, 3})
	resp.Send(rr)

	assert.Equal(t, http.StatusPartialContent, rr.Code)

	var out []int
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, out)
}

func TestStatusMultiStatus(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusMultiStatus("multi")
	resp.Send(rr)

	assert.Equal(t, http.StatusMultiStatus, rr.Code)

	var out string
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, "multi", out)
}

func TestStatusAlreadyReported(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusAlreadyReported(map[string]bool{"done": true})
	resp.Send(rr)

	assert.Equal(t, http.StatusAlreadyReported, rr.Code)

	var out map[string]bool
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.True(t, out["done"])
}

func TestStatusIMUsed(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusIMUsed("im-used")
	resp.SetHeader("cabane", "123")
	resp.Send(rr)

	assert.Equal(t, http.StatusIMUsed, rr.Code)

	var out string
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, "im-used", out)
	assert.Equal(t, "123", rr.Header().Get("cabane"))
}
