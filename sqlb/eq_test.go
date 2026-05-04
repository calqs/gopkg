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
	wb := Where().Eq("id", `"1"`)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id = \"1\"", query)
	assert.Equal(t, []any{}, values)
}

func TestGt(t *testing.T) {
	wb := Where().Gt("id", 1)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id > $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestGte(t *testing.T) {
	wb := Where().Gte("id", 1).Or().Gte("name", `"test"`)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id >= $1 OR name >= \"test\"", query)
	assert.Equal(t, []any{1}, values)
}

func TestGtWithLiteralString(t *testing.T) {
	wb := Where().Gt("id", `"1"`).And().Gt("name", `"test"`)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id > \"1\" AND name > \"test\"", query)
	assert.Equal(t, []any{}, values)
}

func TestLt(t *testing.T) {
	wb := Where().Lt("id", 1)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id < $1", query)
	assert.Equal(t, []any{1}, values)
}

func TestLte(t *testing.T) {
	wb := Where().Lte("id", 1).And().Lte("name", `"test"`)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id <= $1 AND name <= \"test\"", query)
	assert.Equal(t, []any{1}, values)
}

func TestLtWithLiteralString(t *testing.T) {
	wb := Where().Lt("id", `"1"`).And().Lt("name", `"test"`)
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "WHERE id < \"1\" AND name < \"test\"", query)
	assert.Equal(t, []any{}, values)
}
