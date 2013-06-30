package root

import (
	"github.com/Arrow/nm/util"
)

type Newton struct {
	Max_Iter int
	Tol      float64
	fn       util.FuncDiffer
}

func NewNewton(f util.FuncDiffer) *Newton {
	n := new(Newton)
	n.Max_Iter = max_Iter
	n.Tol = tol
	return n
}

func (n *Newton) Compute(x0 float64) (x float64) {
	f := n.fn.Function()
	fp := n.fn.Diff()
	p0 := x0
	p := p0
	for i := 0; i < n.Max_Iter; i++ {
		p = p0 - f(p0)/fp(p0)
		if util.Tolerance(p, p0, n.Tol) {
			return p
		}
	}
	p0 = p
	return p
}
