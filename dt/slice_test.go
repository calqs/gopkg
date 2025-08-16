package dt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I_Can_Append_Variable_Values_Into_A_Slice(t *testing.T) {
	trial := []string{"salut", "les", "kids"}
	var salut, les, kids string

	AppendValues(trial, &salut, &les, &kids)
	assert.Equal(t, trial[0], salut)
	assert.Equal(t, trial[1], les)
	assert.Equal(t, trial[2], kids)
}

func TestMatchAllFunc_Ints(t *testing.T) {
	t.Run("all match", func(t *testing.T) {
		ints := []int{2, 4, 6, 8}
		assert.True(t, MatchAllFunc(ints, func(n int) bool { return n%2 == 0 }))
	})

	t.Run("some do not match", func(t *testing.T) {
		ints := []int{2, 3, 4}
		assert.False(t, MatchAllFunc(ints, func(n int) bool { return n%2 == 0 }))
	})

	t.Run("empty slice is vacuously true", func(t *testing.T) {
		var ints []int
		assert.True(t, MatchAllFunc(ints, func(n int) bool { return n > 0 }))
	})
}

func TestMatchAll(t *testing.T) {
	t.Run("all equal strings", func(t *testing.T) {
		ss := []string{"go", "go", "go"}
		assert.True(t, MatchAll(ss, "go"))
	})

	t.Run("not all equal strings", func(t *testing.T) {
		ss := []string{"go", "rust", "go"}
		assert.False(t, MatchAll(ss, "go"))
	})

	t.Run("empty slice is true", func(t *testing.T) {
		var ss []string
		assert.True(t, MatchAll(ss, "anything"))
	})
}

func TestAnyFunc_Found_FirstMatch(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	users := []user{
		{ID: 1, Name: "Ana"},
		{ID: 2, Name: "Bob"},
		{ID: 2, Name: "Bobby"},
		{ID: 3, Name: "Cid"},
	}

	got, err := MatchAnyFunc(users, func(u user) bool { return u.ID == 2 })
	assert.NoError(t, err)
	assert.Equal(t, got.Name, "Bob")
}

func TestAnyFunc_NotFound_ReturnsZeroAndError(t *testing.T) {
	// Using a non-comparable element type to also exercise the `any` constraint:
	type rec struct{ X int }
	recs := []rec{{1}, {2}, {3}}

	got, err := MatchAnyFunc(recs, func(r rec) bool { return r.X == 42 })
	assert.ErrorIs(t, err, ErrAnyCouldNotFind)
	assert.Equal(t, got, (rec{}))
}

func TestAny_Found(t *testing.T) {
	ints := []int{10, 20, 30}
	got, err := MatchAny(ints, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 20 {
		t.Fatalf("expected 20, got %d", got)
	}
}

func TestAny_NotFound(t *testing.T) {
	ints := []int{1, 2, 3}
	got, err := MatchAny(ints, 99)
	if !errors.Is(err, ErrAnyCouldNotFind) {
		t.Fatalf("expected ErrAnyCouldNotFind, got %v", err)
	}
	if got != 0 { // zero value for int
		t.Fatalf("expected zero value 0, got %d", got)
	}
}

func TestAnyFunc_WithNonComparableElementType(t *testing.T) {
	// Slice of maps (element type is non-comparable), match by a property.
	s := []map[string]int{
		{"a": 1},
		{"b": 2},
		{"c": 3},
	}
	got, err := MatchAnyFunc(s, func(m map[string]int) bool {
		_, ok := m["b"]
		return ok
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := got["b"]; !ok {
		t.Fatalf("expected map containing key 'b', got %#v", got)
	}
}
