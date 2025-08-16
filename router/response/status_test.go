package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Dummy transformer for testing
type DummyTransformer struct{}

func (DummyTransformer) Transform(data any) []byte {
	// Simple: return input unchanged
	return data.([]byte)
}

// --- 2xx ---

func TestStatusOk(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusOk[DummyTransformer]([]byte("ok"))
	resp.Send(rr)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, []byte("ok"), rr.Body.Bytes())
}

func TestStatusCreated(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusCreated[DummyTransformer]([]byte("created"))
	resp.Send(rr)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, []byte("created"), rr.Body.Bytes())
}

func TestStatusAccepted(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusAccepted[DummyTransformer]([]byte("accepted"))
	resp.Send(rr)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	assert.Equal(t, []byte("accepted"), rr.Body.Bytes())
}

func TestStatusNonAuthoritativeInfo(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusNonAuthoritativeInfo[DummyTransformer]([]byte("info"))
	resp.Send(rr)

	assert.Equal(t, http.StatusNonAuthoritativeInfo, rr.Code)
	assert.Equal(t, []byte("info"), rr.Body.Bytes())
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
	resp := StatusPartialContent[DummyTransformer]([]byte("part"))
	resp.Send(rr)

	assert.Equal(t, http.StatusPartialContent, rr.Code)
	assert.Equal(t, []byte("part"), rr.Body.Bytes())
}

func TestStatusMultiStatus(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusMultiStatus[DummyTransformer]([]byte("multi"))
	resp.Send(rr)

	assert.Equal(t, http.StatusMultiStatus, rr.Code)
	assert.Equal(t, []byte("multi"), rr.Body.Bytes())
}

func TestStatusAlreadyReported(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusAlreadyReported[DummyTransformer]([]byte("report"))
	resp.Send(rr)

	assert.Equal(t, http.StatusAlreadyReported, rr.Code)
	assert.Equal(t, []byte("report"), rr.Body.Bytes())
}

func TestStatusIMUsed(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusIMUsed[DummyTransformer]([]byte("im"))
	resp.Send(rr)

	assert.Equal(t, http.StatusIMUsed, rr.Code)
	assert.Equal(t, []byte("im"), rr.Body.Bytes())
}

// --- 3xx ---

func TestStatusMovedPermanently(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusMovedPermanently("https://example.com")
	resp.Send(rr)

	assert.Equal(t, http.StatusMovedPermanently, rr.Code)
	assert.Equal(t, "https://example.com", rr.Header().Get("Location"))
}

func TestStatusFound(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusFound("/found")
	resp.Send(rr)

	assert.Equal(t, http.StatusFound, rr.Code)
	assert.Equal(t, "/found", rr.Header().Get("Location"))
}

func TestStatusSeeOther(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusSeeOther("/other")
	resp.Send(rr)

	assert.Equal(t, http.StatusSeeOther, rr.Code)
	assert.Equal(t, "/other", rr.Header().Get("Location"))
}

func TestNotModified(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := NotModified[DummyTransformer]()
	resp.Send(rr)

	assert.Equal(t, http.StatusNotModified, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
}

func TestStatusTemporaryRedirect(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusTemporaryRedirect("/temp")
	resp.Send(rr)

	assert.Equal(t, http.StatusTemporaryRedirect, rr.Code)
	assert.Equal(t, "/temp", rr.Header().Get("Location"))
}

func TestStatusPermanentRedirect(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := StatusPermanentRedirect("/perm")
	resp.Send(rr)

	assert.Equal(t, http.StatusPermanentRedirect, rr.Code)
	assert.Equal(t, "/perm", rr.Header().Get("Location"))
}

// --- 4xx ---

func TestBadRequest(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := BadRequest("bad")
	resp.Send(rr)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUnauthorized(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := Unauthorized("unauth")
	resp.Send(rr)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestForbidden(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := Forbidden("forbid")
	resp.Send(rr)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := NotFound("nope")
	resp.Send(rr)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestMethodNotAllowed(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := MethodNotAllowed("nah")
	resp.Send(rr)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestTeapot(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := Teapot("tea")
	resp.Send(rr)

	assert.Equal(t, http.StatusTeapot, rr.Code)
}

// --- 5xx ---

func TestInternalServerError(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := InternalServerError("oops")
	resp.Send(rr)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestBadGateway(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := BadGateway("gateway")
	resp.Send(rr)

	assert.Equal(t, http.StatusBadGateway, rr.Code)
}

func TestServiceUnavailable(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := ServiceUnavailable("unavail")
	resp.Send(rr)

	assert.Equal(t, http.StatusServiceUnavailable, rr.Code)
}

func TestGatewayTimeout(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := GatewayTimeout("timeout")
	resp.Send(rr)

	assert.Equal(t, http.StatusGatewayTimeout, rr.Code)
}
