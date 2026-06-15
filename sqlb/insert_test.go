package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_ColumnsValues(t *testing.T) {
	ib := Insert("users").Columns("name", "age").Values("Alice", 30)
	query, values, err := ib.BuildSQL()

	assert.Nil(t, err)
	assert.Equal(t, "INSERT INTO users (name, age) VALUES ($1, $2)", query)
	assert.Equal(t, []any{"Alice", 30}, values)
}

func TestInsert_Set(t *testing.T) {
	ib := Insert("users").Set("name", "Bob").Set("age", 25)
	query, values, err := ib.BuildSQL()

	assert.Nil(t, err)
	assert.Equal(t, "INSERT INTO users (name, age) VALUES ($1, $2)", query)
	assert.Equal(t, []any{"Bob", 25}, values)
}

func TestInsert_Mixed(t *testing.T) {
	ib := Insert("users").Columns("name").Values("Charlie").Set("age", 40)
	query, values, err := ib.BuildSQL()

	assert.Nil(t, err)
	assert.Equal(t, "INSERT INTO users (name, age) VALUES ($1, $2)", query)
	assert.Equal(t, []any{"Charlie", 40}, values)
}

func TestInsert_ErrorLengthMismatch(t *testing.T) {
	ib := Insert("users").Columns("name", "age").Values("Alice")
	query, values, err := ib.BuildSQL()

	assert.NotNil(t, err)
	assert.Equal(t, "InsertBuilder: columns length (2) does not match values length (1)", err.Error())
	assert.Empty(t, query)
	assert.Nil(t, values)
}

func TestInsert_NoColumnsNoValues(t *testing.T) {
	ib := Insert("users")
	query, values, err := ib.BuildSQL()

	assert.Nil(t, err)
	assert.Equal(t, "INSERT INTO users VALUES ()", query)
	assert.Empty(t, values) // values is actually nil
}
