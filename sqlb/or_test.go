package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrWithLike(t *testing.T) {
	wb := Where().Eq("id", 1).Or().Like("name", "test", WildcardBoth)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 OR name LIKE '%' || $2 || '%'", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestOrWithILike(t *testing.T) {
	wb := Where().Eq("id", 1).Or().ILike("name", "test", WildcardBoth)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 OR name ILIKE '%' || $2 || '%'", query)
	assert.Equal(t, []any{1, "test"}, values)
}
