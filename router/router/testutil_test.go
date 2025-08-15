package router

import (
	"encoding/json"
	"net/http"
)

type FuncResponse struct {
	data            []byte
	responder       func([]byte, http.ResponseWriter)
	dataTransformer func(any) []byte
}

func (hr *FuncResponse) Send(w http.ResponseWriter) {
	hr.responder(hr.data, w)
}

func (hr *FuncResponse) WithAnyData(anyData any) *FuncResponse {
	data := hr.dataTransformer(anyData)
	return &FuncResponse{data, hr.responder, hr.dataTransformer}
}

func (hr *FuncResponse) WithDataTransformer(dr func(any) []byte) *FuncResponse {
	return &FuncResponse{hr.data, hr.responder, dr}
}

func (hr *FuncResponse) WithResponser(res func([]byte, http.ResponseWriter)) *FuncResponse {
	return &FuncResponse{hr.data, res, hr.dataTransformer}
}

func NewFuncResponse(
	responder func([]byte, http.ResponseWriter),
	dataTr func(any) []byte,
) *FuncResponse {
	return &FuncResponse{[]byte{}, responder, dataTr}
}

func NewJsonResponse(
	responder func([]byte, http.ResponseWriter),
) *FuncResponse {
	return &FuncResponse{[]byte{}, responder, func(a any) []byte {
		res, _ := json.Marshal(a)
		return res
	}}
}

func NewStringResponse(
	responder func([]byte, http.ResponseWriter),
) *FuncResponse {
	return &FuncResponse{[]byte{}, responder, func(a any) []byte {
		return []byte(a.(string))
	}}
}
