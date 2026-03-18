package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLike(t *testing.T) {
	wb := Where().Like("name", "test", WildcardStart)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE name LIKE '%' || $1", query)
	assert.Equal(t, []any{"test"}, values)
}

func TestILike(t *testing.T) {
	wb := Where().ILike("name", "test", WildcardEnd)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE name ILIKE $1 || '%'", query)
	assert.Equal(t, []any{"test"}, values)
}

func TestLikeChain(t *testing.T) {
	wb := Where().Like("name", "test", WildcardBoth).And().ILike("name2", "test2%", WildcardStart)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE name LIKE '%' || $1 || '%' AND name2 ILIKE '%' || $2", query)
	assert.Equal(t, []any{"test", "test2%"}, values)
}
