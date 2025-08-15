package response

import "encoding/json"

type Transformer interface {
	Transform(any) []byte
}

type JsonTransformer struct{}

func (JsonTransformer) Transform(data any) ([]byte, error) {
	return json.Marshal(data)
}
