package jsonresponse

import (
	"net/http"

	"github.com/calqs/gopkg/router/response"
)

func Answer(data any, code int) *response.HTTPResponse {
	res, err := response.JsonTransformer{}.Transform(data)
	if err != nil {
		errRes := response.InternalServerError("could not transform result", err)
		return &response.HTTPResponse{
			Data:       []byte(errRes.Message),
			StatusCode: errRes.Code,
		}
	}
	return &response.HTTPResponse{
		Data:       res,
		StatusCode: code,
	}
}

// 2xx – Success Responses

func StatusOk(data any) *response.HTTPResponse {
	return Answer(data, http.StatusOK)
}

func StatusCreated(data any) *response.HTTPResponse {
	return Answer(data, http.StatusCreated)
}

func StatusAccepted(data any) *response.HTTPResponse {
	return Answer(data, http.StatusAccepted)
}

func StatusNonAuthoritativeInfo(data any) *response.HTTPResponse {
	return Answer(data, http.StatusNonAuthoritativeInfo)
}

func StatusNoContent() *response.HTTPResponse {
	return &response.HTTPResponse{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}

func StatusResetContent() *response.HTTPResponse {
	return &response.HTTPResponse{
		Data:       nil,
		StatusCode: http.StatusResetContent,
	}
}

func StatusPartialContent(data any) *response.HTTPResponse {
	return Answer(data, http.StatusPartialContent)
}

func StatusMultiStatus(data any) *response.HTTPResponse {
	return Answer(data, http.StatusMultiStatus)
}

func StatusAlreadyReported(data any) *response.HTTPResponse {
	return Answer(data, http.StatusAlreadyReported)
}

func StatusIMUsed(data any) *response.HTTPResponse {
	return Answer(data, http.StatusIMUsed)
}
