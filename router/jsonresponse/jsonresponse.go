package jsonresponse

import (
	"net/http"

	"github.com/calqs/gopkg/router/response"
)

type JSONResponse response.HTTPResponse

func (hr *JSONResponse) Send(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(hr.StatusCode)
	w.Write(hr.Data)
}

func Answer(data any, code int) *JSONResponse {
	res, err := response.JsonTransformer{}.Transform(data)
	if err != nil {
		errRes := response.InternalServerError("could not transform result", err)
		return &JSONResponse{
			Data:       []byte(errRes.Message),
			StatusCode: errRes.Code,
		}
	}
	return &JSONResponse{
		Data:       res,
		StatusCode: code,
	}
}

// 2xx – Success Responses

func StatusOk(data any) *JSONResponse {
	return Answer(data, http.StatusOK)
}

func StatusCreated(data any) *JSONResponse {
	return Answer(data, http.StatusCreated)
}

func StatusAccepted(data any) *JSONResponse {
	return Answer(data, http.StatusAccepted)
}

func StatusNonAuthoritativeInfo(data any) *JSONResponse {
	return Answer(data, http.StatusNonAuthoritativeInfo)
}

func StatusNoContent() *JSONResponse {
	return &JSONResponse{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}

func StatusResetContent() *JSONResponse {
	return &JSONResponse{
		Data:       nil,
		StatusCode: http.StatusResetContent,
	}
}

func StatusPartialContent(data any) *JSONResponse {
	return Answer(data, http.StatusPartialContent)
}

func StatusMultiStatus(data any) *JSONResponse {
	return Answer(data, http.StatusMultiStatus)
}

func StatusAlreadyReported(data any) *JSONResponse {
	return Answer(data, http.StatusAlreadyReported)
}

func StatusIMUsed(data any) *JSONResponse {
	return Answer(data, http.StatusIMUsed)
}
