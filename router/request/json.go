package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JsonBodyRequest safely turns a http.Request.Body data into a struct
func JsonBodyRequest[T any](req *http.Request) (*T, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: http.Request", ErrNilPointer)
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
