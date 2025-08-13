package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func JsonBodyRequest[T any](req *http.Request) (*T, error) {
	if req == nil {
		return nil, errors.New("nil *http.Request")
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	var entity T
	if err := json.Unmarshal(body, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}
