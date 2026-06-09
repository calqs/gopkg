package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	wb := Where().Eq("id", 1).And().IsNull("deleted_at")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 AND deleted_at IS NULL", query)
	assert.Equal(t, []any{1}, values)
}

func TestAndLike(t *testing.T) {
	wb := Where().Eq("id", 1).And().Like("name", "test", WildcardBoth)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 AND name LIKE '%' || $2 || '%'", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestAndILike(t *testing.T) {
	wb := Where().Eq("id", 1).And().ILike("name", "test", WildcardBoth)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 AND name ILIKE '%' || $2 || '%'", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestAndBlock(t *testing.T) {
	wb := Where().Eq("test", 2).And().IsNull("deleted_at").And()
	wb.AndBlock(Eq("id", 1).Or().IsNull("deleted_at"), Eq("name", "test"))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE test = $1 AND deleted_at IS NULL AND ( id = $2 OR deleted_at IS NULL AND name = $3 )", query)
	assert.Equal(t, []any{2, 1, "test"}, values)
}
