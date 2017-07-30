package framework

import "math"

const floatPrecision = 0.000001

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func FloatEq(a, b float64) bool {
	return math.Abs(a-b) < floatPrecision
}

func FloatLs(a, b float64) bool {
	return a+floatPrecision <= b
}
