package framework

import "math"

const FloatPrecision = 0.000001

func Round(f float64) float64 {
	return math.Floor(f + .5)
}
