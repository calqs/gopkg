package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhereBuilderInsideWhere(t *testing.T) {
	wb := Where(Eq(`"cabane".id`, 1).And().IsNull("deleted_at").And().Eq("name", "test"))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, `WHERE "cabane".id = $1 AND deleted_at IS NULL AND name = $2`, query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestWhereBuilderWithAllNodesAsVaradicParams(t *testing.T) {
	wb := Where(Eq("id", 1), And(), IsNull("deleted_at"), And(), Eq("name", "test"))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 AND deleted_at IS NULL AND name = $2", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestChainWhereBuilder(t *testing.T) {
	wb := Where().Eq("id", 1).And().IsNull("deleted_at").And().Eq("name", "test")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 AND deleted_at IS NULL AND name = $2", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestWhereBuilderChunkedUse(t *testing.T) {
	wb := Where(IsNull("deleted_at").And().Eq("id", 1))
	wb.And().Eq("name", "test")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE deleted_at IS NULL AND id = $1 AND name = $2", query)
	assert.Equal(t, []any{1, "test"}, values)
}

func TestOrBlock(t *testing.T) {
	wb := Where().Eq("test", 2).Or().Eq("test2", 3)
	wb.Or().OrBlock(Eq("id", 1).And().IsNull("deleted_at"), Eq("name", "test"))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE test = $1 OR test2 = $2 OR ( id = $3 AND deleted_at IS NULL OR name = $4 )", query)
	assert.Equal(t, []any{2, 3, 1, "test"}, values)
}
