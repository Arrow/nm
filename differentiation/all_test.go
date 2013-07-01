package differentiation

// Define functions to test differentiation.
// From: J. Oliver, An algorithm for numerical differentiation of a function of one real variable,
// Journal of Computational and Applied Mathematics, Volume 6, Issue 2, June 1980, Pages 145-160, ISSN 0377-0427,
// http://dx.doi.org/10.1016/0771-050X(80)90008-X.
// (http://www.sciencedirect.com/science/article/pii/0771050X8090008X)
// f(x), xmin, xmax, a and b from Table 2
// Added f'(x) from wolfram alpha
//
//	xi = a + i * (b - a) / 9 .... i := 0, 1... 9

import (
	"github.com/Arrow/nm/util"
	"math"
)

type test struct {
	f    util.BasicFuncPair
	a, b float64
}

func (t test) XYs() []util.TestXY {
	xy := make([]util.TestXY, 10)
	for i := 0; i < 10; i++ {
		xy[i].X = t.a + float64(i)*(t.b-t.a)/9
		fp := t.f.Diff()
		xy[i].Y = fp(xy[i].X)
	}
	return xy
}

var diff_test []test

func init() {
	diff_test = make([]test, 13)
	// 1  - f(x) = exp(4x)
	//      f'(x) = 4 exp(4x)
	//	xmin = -3
	//	xmax =  3
	//	a = -1
	//	b =  2
	diff_test[0].f.F = func(x float64) float64 { return math.Exp(4 * x) }
	diff_test[0].f.Fp = func(x float64) float64 { return 4 * math.Exp(4*x) }
	diff_test[0].a, diff_test[0].b = -1, 2

	// 2  - f(x) = exp(x^2)
	//      f'(x) = 2x exp(x^2)
	//	xmin = -3
	//	xmax =  3
	//	a =  0
	//	b =  2
	diff_test[1].f.F = func(x float64) float64 { return math.Exp(math.Pow(x, 2)) }
	diff_test[1].f.Fp = func(x float64) float64 { return 2 * x * math.Exp(math.Pow(x, 2)) }
	diff_test[1].a, diff_test[1].b = 0, 2

	// 3  - f(x) = sqrt(x)
	//      f'(x) = 1/(2 sqrt(x))
	//	xmin =  10E-10
	//	xmax =  1000
	//	a =  0.005
	//	b =  1
	diff_test[2].f.F = func(x float64) float64 { return math.Sqrt(x) }
	diff_test[2].f.Fp = func(x float64) float64 { return 1 / (2 * math.Sqrt(x)) }
	diff_test[2].a, diff_test[2].b = 0.005, 1

	// 4  - f(x) = 1/x
	//      f'(x) = - 1/x^2
	//	xmin =  10E-10
	//	xmax =  1000
	//	a =  0.005
	//	b =  1
	diff_test[3].f.F = func(x float64) float64 { return 1 / x }
	diff_test[3].f.Fp = func(x float64) float64 { return -1 / math.Pow(x, 2) }
	diff_test[3].a, diff_test[3].b = 0.005, 1

	// 5  - f(x) = ln x
	//      f'(x) = 1/x
	//	xmin =  10E-10
	//	xmax =  1000
	//	a =  0.005
	//	b =  1
	diff_test[4].f.F = func(x float64) float64 { return math.Log(x) }
	diff_test[4].f.Fp = func(x float64) float64 { return 1 / x }
	diff_test[4].a, diff_test[4].b = 0.005, 1

	// 6  - f(x) = x^2 ln x
	//      f'(x) = x+2 x log(x)
	//	xmin =  10E-10
	//	xmax =  1000
	//	a =  0.005
	//	b =  1
	diff_test[5].f.F = func(x float64) float64 { return math.Pow(x, 2) * math.Log(x) }
	diff_test[5].f.Fp = func(x float64) float64 { return x + 2*x*math.Log(x) }
	diff_test[5].a, diff_test[5].b = 0.005, 1

	// 7  - f(x) = ln [x + sqrt(1+x^2)]
	//      f'(x) = 1/sqrt(1+x^2)
	//	xmin = -1000
	//	xmax =  1000
	//	a = -2
	//	b =  5
	diff_test[6].f.F = func(x float64) float64 { return math.Log(x + math.Sqrt(1+math.Pow(x, 2))) }
	diff_test[6].f.Fp = func(x float64) float64 { return 1 / math.Sqrt(1+math.Pow(x, 2)) }
	diff_test[6].a, diff_test[6].b = -2, 5

	// 8  - f(x) = sin x
	//      f'(x) = cos x
	//	xmin = -1000
	//	xmax =  1000
	//	a =  0
	//	b =  2
	diff_test[7].f.F = func(x float64) float64 { return math.Sin(x) }
	diff_test[7].f.Fp = func(x float64) float64 { return math.Cos(x) }
	diff_test[7].a, diff_test[7].b = 0, 2

	// 9  - f(x) = sin( x^2 )
	//      f'(x) = 2 x cos(x^2) *Use for Root Finder
	//	xmin = -1000
	//	xmax =  1000
	//	a =  0
	//	b =  2
	diff_test[8].f.F = func(x float64) float64 { return math.Sin(math.Pow(x, 2)) }
	diff_test[8].f.Fp = func(x float64) float64 { return 2 * x * math.Cos(math.Pow(x, 2)) }
	diff_test[8].a, diff_test[8].b = 0, 2

	// 10 - f(x) = x^2 cos(2x)
	//      f'(x) = 2 x (cos(2 x)-x sin(2 x))
	//	xmin = -1000
	//	xmax =  1000
	//	a =  0
	//	b =  1
	diff_test[9].f.F = func(x float64) float64 { return math.Pow(x, 2) * math.Cos(2*x) }
	diff_test[9].f.Fp = func(x float64) float64 { return 2 * x * (math.Cos(2*x) - x*math.Sin(2*x)) }
	diff_test[9].a, diff_test[9].b = 0, 1

	// 11 - f(x) = sin^2(x)
	//      f'(x) = sin (2x)
	//	xmin = -1000
	//	xmax =  1000
	//	a =  0
	//	b =  2
	diff_test[10].f.F = func(x float64) float64 { return math.Pow(math.Sin(x), 2) }
	diff_test[10].f.Fp = func(x float64) float64 { return math.Sin(2 * x) }
	diff_test[10].a, diff_test[10].b = 0, 2

	// 12 - f(x) = cos(ln x)
	//      f'(x) = -(sin(ln(x)))/x
	//	xmin =  10E-10
	//	xmax =  1000
	//	a =  0.005
	//	b =  1
	diff_test[11].f.F = func(x float64) float64 { return math.Cos(math.Log(x)) }
	diff_test[11].f.Fp = func(x float64) float64 { return -(math.Sin(math.Log(x))) / x }
	diff_test[11].a, diff_test[11].b = 0.005, 1

	// 13 - f(x) = (x^2 + 0.01)^-1
	//      f'(x) = -(2. x)/(0.01+x^2)^2
	//	xmin = -1000
	//	xmax =  1000
	//	a =  0
	//	b =  0.5
	diff_test[12].f.F = func(x float64) float64 { return 1 / (math.Pow(x, 2) + 0.01) }
	diff_test[12].f.Fp = func(x float64) float64 { return -2 * x / math.Pow(x*x+0.01, 2) }
	diff_test[12].a, diff_test[12].b = 0, 0.5
}

/*
var diff_test = test{
	util.BasicFunc(func(x float64) float64 { return math.Pow(x, 2) }),
	[]util.TestXY{
		{-5, -10},
		{-2, -4},
		{0, 0},
		{2, 4},
		{5, 10},
	},
}*/
