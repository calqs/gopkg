package router

import (
	"encoding/json"
	"net/http"

	"github.com/calqs/gopkg/router/response"
)

type FuncResponse struct {
	response.ResponseHeaders
	data            []byte
	responder       func([]byte, http.ResponseWriter)
	dataTransformer func(any) []byte
}

func (hr *FuncResponse) Send(w http.ResponseWriter) {
	hr.responder(hr.data, w)
}

func (hr *FuncResponse) WithAnyData(anyData any) *FuncResponse {
	data := hr.dataTransformer(anyData)
	return &FuncResponse{data: data, responder: hr.responder, dataTransformer: hr.dataTransformer}
}

func (hr *FuncResponse) WithDataTransformer(dr func(any) []byte) *FuncResponse {
	return &FuncResponse{data: hr.data, responder: hr.responder, dataTransformer: dr}
}

func (hr *FuncResponse) WithResponser(res func([]byte, http.ResponseWriter)) *FuncResponse {
	return &FuncResponse{data: hr.data, responder: res, dataTransformer: hr.dataTransformer}
}

func NewFuncResponse(
	responder func([]byte, http.ResponseWriter),
	dataTr func(any) []byte,
) *FuncResponse {
	return &FuncResponse{data: []byte{}, responder: responder, dataTransformer: dataTr}
}

func NewJsonResponse(
	responder func([]byte, http.ResponseWriter),
) *FuncResponse {
	return &FuncResponse{data: []byte{}, responder: responder, dataTransformer: func(a any) []byte {
		res, _ := json.Marshal(a)
		return res
	}}
}

func NewStringResponse(
	responder func([]byte, http.ResponseWriter),
) *FuncResponse {
	return &FuncResponse{data: []byte{}, responder: responder, dataTransformer: func(a any) []byte {
		return []byte(a.(string))
	}}
}
