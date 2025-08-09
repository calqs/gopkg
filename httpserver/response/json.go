package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponseWriter(data *any, w http.ResponseWriter) {
	res, err := json.Marshal(data)
	if err != nil {
		InternalServerError("Could not marshal json matters").Write(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
