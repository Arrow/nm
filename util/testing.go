package util

import (
	"math"
)

func Tolerance(value float64, target float64, tol float64) bool {
	if (math.Abs(value - target) > tol) {
		return false
	} else {
		return true
	}
}
