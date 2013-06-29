package differentiation

import (
	"github.com/Arrow/nm/util"
)

func FivePoint(fn util.Functioner, x float64) (fp float64) {
	h := StepSize
	f := fn.Function()
	return ((f(x-2*h) - 8*f(x-h) + 8*f(x+h) - f(x+2*h)) / (12 * h))
}
