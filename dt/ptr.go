package dt

// Ptr is a generic one-liner for making a pointer
// ex: Ptr("Im just a poor boy")
func Ptr[T any](value T) *T {
	return &value
}

// PtrNilOnEmpty is a generic one-liner for making a pointer, or nil if empty
// ex: Ptr("") => nil
func PtrNilOnEmpty[T any](value T) *T {
	switch t := any(value).(type) {
	case string:
		if t != "" {
			return &value
		}
	}
	return nil
}

// Deref is a generic one-liner for safely dereferencing a pointer
// ex: Deref(&test{}) => test{}
// ex: Deref[test](nil) => test{}
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var nope T
		return nope
	}
	return *ptr
}

// Deref is a generic one-liner for safely dereferencing a pointer,
// and choosing a default behavior in case dereferencing is not possible
// ex: Deref(test{"a"}, test{"b"}) => test{"a"}
// ex: Deref(nil, test{"b"}) => test{"b"}
func DerefOr[T any](ptr *T, orElse T) T {
	if ptr == nil {
		return orElse
	}
	return *ptr
}
