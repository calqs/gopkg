package ptr

func Ptr[T any](value T) *T {
	return &value
}

func PtrNilOnEmpty[T any](value T) *T {
	switch t := any(value).(type) {
	case string:
		if t != "" {
			return &value
		}
	}
	return nil
}

func Deref[T any](ptr *T) T {
	if ptr == nil {
		var nope T
		return nope
	}
	return *ptr
}

func DerefOr[T any](ptr *T, orElse T) T {
	if ptr == nil {
		return orElse
	}
	return *ptr
}
