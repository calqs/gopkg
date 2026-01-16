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
	if req.Body == nil {
		return nil, nil
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
		switch err.(type) {
		case *json.UnmarshalTypeError:
			return nil, fmt.Errorf("%w: %w", ErrPayloadWrongFieldType, err)
		case *json.SyntaxError:
			return nil, fmt.Errorf("%w: %w", ErrPayloadWrongShape, err)
		}
		return nil, err
	}
	return &entity, nil
}
