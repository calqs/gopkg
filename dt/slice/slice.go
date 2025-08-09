package slice

import "errors"

var (
	ErrAnyCouldNotFind = errors.New("unable to find any matching item")
)

func MapVars[T any](slice []T, into ...*T) {
	ls := len(slice)
	li := len(into)
	for i := 0; i < ls && i < li; i++ {
		*into[i] = slice[i]
	}
}

func MatchAllFunc[ItemT comparable](s []ItemT, cond func(ItemT) bool) bool {
	for _, item := range s {
		if !cond(item) {
			return false
		}
	}
	return true
}

func MatchAll[ItemT comparable](s []ItemT, value ItemT) bool {
	return MatchAllFunc(s, func(item ItemT) bool { return item == value })
}

func AnyFunc[ItemT any](slice []ItemT, match func(ItemT) bool) (ItemT, error) {
	for _, item := range slice {
		if match(item) {
			return item, nil
		}
	}
	var dummy ItemT
	return dummy, ErrAnyCouldNotFind
}

func Any[ItemT comparable](slice []ItemT, value ItemT) (ItemT, error) {
	return AnyFunc(slice, func(it ItemT) bool { return it == value })
}
