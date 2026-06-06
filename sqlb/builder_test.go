package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	wb := Select("id", "SUM(amount)").From("expenses").GroupBy("id")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT id, SUM(amount) FROM expenses GROUP BY id", query)
	assert.Empty(t, values)
}

func TestGroupByMultipleCols(t *testing.T) {
	wb := Select("category", "status", "COUNT(1)").From("tasks").GroupBy("category", "status")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT category, status, COUNT(1) FROM tasks GROUP BY category, status", query)
	assert.Empty(t, values)
}

func TestGroupByChained(t *testing.T) {
	wb := Select("year", "month", "SUM(amount)").From("sales").GroupBy("year").GroupBy("month")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT year, month, SUM(amount) FROM sales GROUP BY year, month", query)
	assert.Empty(t, values)
}

func TestGroupByWithWhere(t *testing.T) {
	wb := Select("status", "COUNT(id)").From("users").Where(Eq("active", true)).GroupBy("status")
	query, values, err := wb.BuildSQL()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT status, COUNT(id) FROM users WHERE active = $1 GROUP BY status", query)
	assert.Equal(t, []any{true}, values)
}
