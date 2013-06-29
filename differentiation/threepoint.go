package differentiation

import (
	"github.com/Arrow/nm/util"
)

func ThreePoint(fn util.Functioner, x float64) (fp float64) {
	h := StepSize
	f := fn.Function()
	return ((f(x+h) - f(x-h)) / (2 * h))
}
