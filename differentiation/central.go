package differentiation

import (
	"github.com/Arrow/nm/util"
	"math"
)

type Central struct {
	StepSize float64
	fn       util.Functioner
	er       float64
}

func NewCentral(fn util.Functioner) *Central {
	c := new(Central)
	c.StepSize = StepSize
	c.fn = fn
	return c
}

func (c *Central) Compute(x float64) (fp float64) {
	h := c.StepSize
	f := c.fn.Function()
	fp = ((f(x-2*h) - 8*f(x-h) + 8*f(x+h) - f(x+2*h)) / (12 * h))
	fp3 := ((f(x+h) - f(x-h)) / (2 * h))
	c.er = math.Abs(fp - fp3)
	return fp
}
