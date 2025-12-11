package dt

import "math"

func RoundFloat(f float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return math.Round(f*p) / p
}
