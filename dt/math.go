package dt

import "math"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Sum[T any, ResT Number](a []T, fn func(T) ResT) ResT {
	var res ResT
	for _, v := range a {
		res += fn(v)
	}
	return res
}

func RoundFloat(f float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return math.Round(f*p) / p
}
