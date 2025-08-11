package router

import (
	"encoding/json"
	"net/http"

	"github.com/calqs/gopkg/router/response"
)

func JsonResponseWriter(data any, w http.ResponseWriter) {
	res, err := json.Marshal(data)
	if err != nil {
		response.InternalServerError("Could not marshal json matters").Write(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
