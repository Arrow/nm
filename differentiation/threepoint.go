package differentiation

import (
	"github.com/Arrow/nm/util"
)

func ThreePoint(f util.Function, x float64) (fp float64) {
	h := StepSize
	return ((f(x+h)-f(x-h))/(2*h))
}
