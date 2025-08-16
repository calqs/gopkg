package response

import "net/http"

// 2xx – Success Responses

func StatusOk[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusOK,
	}
}

func StatusCreated[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusCreated,
	}
}

func StatusAccepted[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusAccepted,
	}
}

func StatusNonAuthoritativeInfo[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusNonAuthoritativeInfo,
	}
}

func StatusNoContent() *HTTPResponse {
	return &HTTPResponse{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}

func StatusResetContent() *HTTPResponse {
	return &HTTPResponse{
		Data:       nil,
		StatusCode: http.StatusResetContent,
	}
}

func StatusPartialContent[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusPartialContent,
	}
}

func StatusMultiStatus[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusMultiStatus,
	}
}

func StatusAlreadyReported[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusAlreadyReported,
	}
}

func StatusIMUsed[TransformerT Transformer](data any) *HTTPResponse {
	var transformer TransformerT
	return &HTTPResponse{
		Data:       transformer.Transform(data),
		StatusCode: http.StatusIMUsed,
	}
}

// 3xx – Redirection Responses

func StatusMultipleChoices(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusMultipleChoices,
	}
}

func StatusMovedPermanently(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusMovedPermanently,
	}
}

func StatusFound(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusFound,
	}
}

func StatusSeeOther(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusSeeOther,
	}
}

func NotModified[TransformerT Transformer]() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusNotModified,
	}
}

// Special: 304 Not Modified → no Location, no body (fits better in HTTPResponse with nil data)

func StatusUseProxy(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusUseProxy,
	}
}

func StatusTemporaryRedirect(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusTemporaryRedirect,
	}
}

func StatusPermanentRedirect(location string) *RedirectResponse {
	return &RedirectResponse{
		location:   location,
		StatusCode: http.StatusPermanentRedirect,
	}
}

// 4xx – Client Errors

func BadRequest(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusBadRequest, msg, errs...)
}

func Unauthorized(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusUnauthorized, msg, errs...)
}

func Forbidden(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusForbidden, msg, errs...)
}

func NotFound(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusNotFound, msg, errs...)
}

func MethodNotAllowed(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusMethodNotAllowed, msg, errs...)
}

func RequestTimeout(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusRequestTimeout, msg, errs...)
}

func Conflict(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusConflict, msg, errs...)
}

func Gone(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusGone, msg, errs...)
}

func Teapot(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusTeapot, msg, errs...)
}

func TooManyRequests(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusTooManyRequests, msg, errs...)
}

func UnprocessableEntity(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusUnprocessableEntity, msg, errs...)
}

// 5xx – Server Errors

func InternalServerError(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusInternalServerError, msg, errs...)
}

func NotImplemented(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusNotImplemented, msg, errs...)
}

func BadGateway(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusBadGateway, msg, errs...)
}

func ServiceUnavailable(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusServiceUnavailable, msg, errs...)
}

func GatewayTimeout(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusGatewayTimeout, msg, errs...)
}

func HTTPVersionNotSupported(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusHTTPVersionNotSupported, msg, errs...)
}

func VariantAlsoNegotiates(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusVariantAlsoNegotiates, msg, errs...)
}

func InsufficientStorage(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusInsufficientStorage, msg, errs...)
}

func LoopDetected(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusLoopDetected, msg, errs...)
}

func NotExtended(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusNotExtended, msg, errs...)
}

func NetworkAuthenticationRequired(msg string, errs ...error) *HTTPError {
	return NewError(http.StatusNetworkAuthenticationRequired, msg, errs...)
}
