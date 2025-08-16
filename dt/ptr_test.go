package dt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestICanMakeAPtrOrNilOnEmptyValue(t *testing.T) {
	assert.Nil(t, PtrNilOnEmpty(""))
	assert.Equal(t, "a", *PtrNilOnEmpty("a"))
}

func TestICanDerefPointers(t *testing.T) {
	type fakeStruct struct {
		cabane string
	}
	assert.Equal(t, "cabane1", Deref(Ptr("cabane1")))
	assert.Equal(t, fakeStruct{}, Deref(&fakeStruct{}))
	assert.Equal(t, fakeStruct{}, Deref[fakeStruct](nil))
	assert.Equal(t, 123, DerefOr(nil, 123))
	var test1 *fakeStruct
	assert.Equal(t, fakeStruct{}, DerefOr(test1, fakeStruct{}))
	assert.Equal(t, fakeStruct{"123"}, DerefOr(&fakeStruct{"123"}, fakeStruct{}))
}
