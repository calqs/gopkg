package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonTransformer_Transform_Success(t *testing.T) {
	got, err := JsonTransformer{}.Transform(struct{ Name string }{"Alice"})
	assert.NoError(t, err)
	var decoded any
	assert.NoError(t, json.Unmarshal(got, &decoded))
}

func TestJsonTransformer_Transform_Error(t *testing.T) {
	_, err := JsonTransformer{}.Transform(make(chan int))
	assert.Error(t, err)
	_, err = JsonTransformer{}.Transform(func() {})
	assert.Error(t, err)
}
