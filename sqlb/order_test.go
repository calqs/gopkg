package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	wb := Where(Eq("id", 1)).OrderBy("id").OrderDir(Asc)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 ORDER BY id ASC", query)
	assert.Equal(t, []any{1}, values)
}

func TestOrderDesc(t *testing.T) {
	wb := Where(Eq("id", 1)).OrderBy("id").OrderDir(Desc)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = $1 ORDER BY id DESC", query)
	assert.Equal(t, []any{1}, values)
}
