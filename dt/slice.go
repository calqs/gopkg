package dt

// AppendValues is a generic function for inserting values into a slice
// ex: MapVars(res, "I", "need", "no", "sympathy") => []string{"I", "need", "no", "sympathy"}
func AppendValues[T any](slice []T, into ...*T) {
	ls := len(slice)
	li := len(into)
	for i := 0; i < ls && i < li; i++ {
		*into[i] = slice[i]
	}
}

// MatchAllFunc is a generic function verifying each item matches a predicate
// ex1: MatchAllFunc([]int{2, 4, 6}, func(item int) bool {return item%2 == 0}) => true
// ex2: MatchAllFunc([]int{2, 4, 5}, func(item int) bool {return item%2 == 0}) => false
func MatchAllFunc[ItemT comparable](s []ItemT, cond func(ItemT) bool) bool {
	for _, item := range s {
		if !cond(item) {
			return false
		}
	}
	return true
}

// MatchAll is a generic function verifying each item matches a value
// ex1: MatchAll([]string{"go", "go"}, "go") => true
// ex1: MatchAll([]string{"johnny", "b.", "goode"}, "goode") => false
func MatchAll[ItemT comparable](s []ItemT, value ItemT) bool {
	return MatchAllFunc(s, func(item ItemT) bool { return item == value })
}

// MatchAnyFunc is a generic function trying to find at least 1 element matching a predicate
// ex1: MatchAnyFunc([]int{1, 2, 5}, func(item int) bool {return item % 5 == 0}) => true
// ex2: MatchAnyFunc([]int{1, 2, 6}, func(item int) bool {return item % 5 == 0}) => false
func MatchAnyFunc[ItemT any](slice []ItemT, match func(ItemT) bool) (ItemT, error) {
	for _, item := range slice {
		if match(item) {
			return item, nil
		}
	}
	var dummy ItemT
	return dummy, ErrAnyCouldNotFind
}

// MatchAny is a generic function trying to find at least 1 element matching a value
// ex1: MatchAny([]string{"got", "me", "on", "my", "knees"}, "layla") => false
// ex2: MatchAny([]string{"like", "a", "fool"}, "fool") => true
func MatchAny[ItemT comparable](slice []ItemT, value ItemT) (ItemT, error) {
	return MatchAnyFunc(slice, func(it ItemT) bool { return it == value })
}

func SliceTransform[FromT, ToT any](from []FromT, transformer func(FromT) ToT) []ToT {
	res := make([]ToT, len(from))
	for k, val := range from {
		res[k] = transformer(val)
	}
	return res
}

func SliceFilterFunc[ItemT any](from []ItemT, filter func(ItemT) bool) []ItemT {
	res := make([]ItemT, 0, len(from))
	for _, val := range from {
		if filter(val) {
			res = append(res, val)
		}
	}
	return res
}
