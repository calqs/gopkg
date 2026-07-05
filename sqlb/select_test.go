package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	wb := Select("id", "name").From("users").Where(Eq("id", 1))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT id, name FROM users WHERE id = $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestFromSelectWhere(t *testing.T) {
	wb := From("users").Select("id", "name").Where(Eq("id", 1))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT id, name FROM users WHERE id = $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestFromWhereSelect(t *testing.T) {
	wb := From("users").Where(Eq("id", 1)).Select("id").Select("name")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT name FROM users WHERE id = $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestSelectAll(t *testing.T) {
	wb := Select().From("users").Where(Eq("id", 1))
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT * FROM users WHERE id = $1", query)
	assert.Equal(t, []any{1}, values)
}
