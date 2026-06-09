package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoid(t *testing.T) {
	v := Void()
	assert.NotNil(t, v)
	// Gte
	assert.Equal(t, Gte("id", 1), v.Gte("id", 1))
	// Gt
	assert.Equal(t, Gt("id", 1), v.Gt("id", 1))
	// Lt
	assert.Equal(t, Lt("id", 1), v.Lt("id", 1))
	// Lte
	assert.Equal(t, Lte("id", 1), v.Lte("id", 1))
	// Eq
	assert.Equal(t, Eq("id", 1).Or().Eq("id", 2), v.Eq("id", 1).Or().Eq("id", 2))
	// In
	assert.Equal(t, In("id", 1, 2), v.In("id", 1, 2))
	// IsNull
	assert.Equal(t, IsNull("id"), v.IsNull("id"))
	// NotNull
	assert.Equal(t, NotNull("id"), v.NotNull("id"))
	// And
	assert.Equal(t, And(), v.And())
	// Or
	assert.Equal(t, Or(), v.Or())
	// AndBlock
	assert.Equal(t, AndBlock(Eq("id", 1)), v.AndBlock(Eq("id", 1)))
	// OrBlock
	assert.Equal(t, OrBlock(Eq("id", 1)), v.OrBlock(Eq("id", 1)))
}
