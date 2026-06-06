package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderClone(t *testing.T) {
	wb1 := Select("id", "name").From("users").Where(Eq("status", "active")).And().In("role", "admin", "moderator")
	wb1.And().Like("email", "test", WildcardEnd)

	wb2 := wb1.Clone()

	// Modify wb2 and ensure wb1 is not affected
	wb2.And().Eq("deleted", false)
	wb2.Select("id", "name", "email")
	wb2.OrderBy("id").OrderDir(Desc)

	q1, v1, err1 := wb1.BuildSQL()
	assert.NoError(t, err1)
	assert.Equal(t, "SELECT id, name FROM users WHERE status = $1 AND role IN ($2, $3) AND email LIKE $4 || '%'", q1)
	assert.Equal(t, []any{"active", "admin", "moderator", "test"}, v1)

	q2, v2, err2 := wb2.BuildSQL()
	assert.NoError(t, err2)
	assert.Equal(t, "SELECT id, name, email FROM users WHERE status = $1 AND role IN ($2, $3) AND email LIKE $4 || '%' AND deleted = $5 ORDER BY id DESC", q2)
	assert.Equal(t, []any{"active", "admin", "moderator", "test", false}, v2)
}

func TestCloneJoin(t *testing.T) {
	wb1 := Select("users.id").From("users").LeftJoin("profiles", Eq("users.id", "profiles.user_id")).Where(Eq("users.id", 1))
	wb2 := wb1.Clone()

	wb2.Join("settings", Eq("users.id", "settings.user_id"))

	q1, _, err1 := wb1.BuildSQL()
	assert.NoError(t, err1)
	assert.Equal(t, "SELECT users.id FROM users LEFT JOIN profiles ON users.id = profiles.user_id WHERE users.id = $1", q1)

	q2, _, err2 := wb2.BuildSQL()
	assert.NoError(t, err2)
	assert.Equal(t, "SELECT users.id FROM users LEFT JOIN profiles ON users.id = profiles.user_id INNER JOIN settings ON users.id = settings.user_id WHERE users.id = $1", q2)
}
