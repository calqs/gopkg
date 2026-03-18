package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	wb := Where().Eq("id", 1)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestEqWithLiteralString(t *testing.T) {
	wb := Where().Eq("id", "'1'")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = '1'", query)
	assert.Equal(t, []any{}, values)
}
