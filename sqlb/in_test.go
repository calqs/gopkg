package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInNode(t *testing.T) {
	node := In("id", 1, 2, 3)
	sql, args := node.ToSQL(1)
	assert.Equal(t, "id IN ($1, $2, $3)", sql)
	assert.Equal(t, []any{1, 2, 3}, args)
}

func TestInNodeWithAnd(t *testing.T) {
	wb := Where().In("id", 1, 2, 3).And().Eq("name", "test")
	sql, args, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id IN ($1, $2, $3) AND name = $4", sql)
	assert.Equal(t, []any{1, 2, 3, "test"}, args)
}
