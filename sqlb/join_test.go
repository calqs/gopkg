package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	wb := Select("id", "name").
		From("users").
		Join("orders", Eq("users.id", "orders.user_id")).
		LeftJoin("cabane", Eq("users.id", "cabane.user_id")).
		Where().
		NotNull("test")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT id, name FROM users INNER JOIN orders ON users.id = orders.user_id LEFT JOIN cabane ON users.id = cabane.user_id WHERE test IS NOT NULL", query)
	assert.Equal(t, []any{}, values)
}
