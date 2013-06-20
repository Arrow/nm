package root

import (
	"github.com/Arrow/nm/util"
	derive "github.com/Arrow/nm/differentiation"
)

func Newton(f util.Function, x0 float64) (x float64) {
	p0 := x0
	p  := p0
	for i := 0; i < Max_Iter; i++ {
		p = p0 - f(p0) / derive.FivePoint(f, p0)
		if util.Tolerance(p, p0, tol) { return p }
	}
	p0 = p
	return p
}
