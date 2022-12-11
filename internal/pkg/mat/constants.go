package mat

import "math"

const EPSILON = 0.00001

func AlmostEqualFloat64(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
