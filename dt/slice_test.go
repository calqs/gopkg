package dt

import (
	"strconv"
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

	got := MatchAnyFunc(users, func(u user) bool { return u.ID == 2 })
	assert.NotNil(t, got)
	assert.Equal(t, got.Name, "Bob")
}

func TestAnyFunc_NotFound_ReturnsZeroAndError(t *testing.T) {
	// Using a non-comparable element type to also exercise the `any` constraint:
	type rec struct{ X int }
	recs := []rec{{1}, {2}, {3}}

	got := MatchAnyFunc(recs, func(r rec) bool { return r.X == 42 })
	assert.Nil(t, got)
}

func TestAny_Found(t *testing.T) {
	ints := []int{10, 20, 30}
	got := MatchAny(ints, 20)
	if *got != 20 {
		t.Fatalf("expected 20, got %d", got)
	}
}

func TestAny_NotFound(t *testing.T) {
	ints := []int{1, 2, 3}
	got := MatchAny(ints, 99)
	assert.Nil(t, got)
}

func TestAnyFunc_WithNonComparableElementType(t *testing.T) {
	// Slice of maps (element type is non-comparable), match by a property.
	s := []map[string]int{
		{"a": 1},
		{"b": 2},
		{"c": 3},
	}
	got := MatchAnyFunc(s, func(m map[string]int) bool {
		_, ok := m["b"]
		return ok
	})
	assert.NotNil(t, got)
	if _, ok := (*got)["b"]; !ok {
		t.Fatalf("expected map containing key 'b', got %#v", got)
	}
}

func Test_I_Can_Transform_A_Slice(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, SliceTransform([]int{1, 2, 3}, func(v int) string { return strconv.Itoa(v) }))
	assert.Equal(t, []string{}, SliceTransform([]int{}, func(v int) string { return strconv.Itoa(v) }))
}

func Test_I_Can_Filter_A_Slice(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6, 8}, SliceFilterFunc([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(it int) bool { return it%2 == 0 }))
	assert.Equal(t, []int{}, SliceFilterFunc([]int{}, func(it int) bool { return it%2 == 0 }))
}

func Test_I_Can_Match_Slices(t *testing.T) {
	type testT struct {
		a, b int
	}
	assert.True(t, SlicesMatch([]int{1, 2, 3}, []int{1, 2, 3}, func(a, b int) bool { return a == b }))
	assert.False(t, SlicesMatch([]int{1, 2, 3}, []int{1, 2, 4}, func(a, b int) bool { return a == b }))
	assert.False(t, SlicesMatch([]int{1, 2, 3}, []int{1, 2, 3, 4}, func(a, b int) bool { return a == b }))
	assert.True(t, SlicesMatch([]testT{{1, 2}, {3, 4}}, []testT{{1, 2}, {3, 4}}, func(a, b testT) bool { return a.a == b.a && a.b == b.b }))
}

func TestSortEqual(t *testing.T) {
	assert.True(t, SortEqual([]int{3, 1, 2}, []int{1, 2, 3}))
	assert.False(t, SortEqual([]int{3, 1, 2}, []int{1, 2, 4}))
	assert.False(t, SortEqual([]int{3, 1, 2}, []int{1, 2}))
}

func TestUnique(t *testing.T) {
	assert.EqualValues(t, []int{1, 2, 3}, Unique([]int{1, 2, 3, 1, 2, 3}))
	assert.EqualValues(t, []int{}, Unique([]int{}))
}

func TestSortEqualFunc(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	users_a := []user{{1, "a"}, {2, "b"}}
	users_b := []user{{2, "b"}, {1, "a"}}
	assert.True(t, SortEqualFunc(users_a, users_b, func(a, b user) int { return a.ID - b.ID }))
}

func TestMergeReplace(t *testing.T) {
	type budget struct {
		ID   int
		Name string
	}
	base := []budget{{1, "a"}, {2, "b"}}
	toAdd := []budget{{2, "B"}, {3, "c"}}
	merged := MergeReplace(base, toAdd, func(a, b budget) bool {
		return a.ID == b.ID
	})
	assert.Equal(t, []budget{{1, "a"}, {2, "B"}, {3, "c"}}, merged)
}
