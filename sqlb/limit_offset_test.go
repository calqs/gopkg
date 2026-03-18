package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitOffset(t *testing.T) {
	wb := Where().Eq("id", 1).Limit(10).Offset(20)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 LIMIT 10 OFFSET 20", query)
	assert.Equal(t, []any{1}, values)
}
