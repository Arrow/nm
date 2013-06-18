package differentiation

import (
	"math"
	"testing"
	"github.com/Arrow/nm/util"
)

const (
	tol = 0.0000001
	h = 0.5
)

type TestXY struct {
	x, y float64
}

var diff_test = []TestXY{
	{-5, -10},
	{-2, -4},
	{0, 0},
	{2, 4},
	{5, 10},
}

func testFunc(x float64) float64 {
	return math.Pow(x, 2)
}

func TestThreePoint(t *testing.T) {
	f := testFunc
	for _, c := range diff_test {
		if(!util.Tolerance(ThreePoint(f, c.x, h), c.y, tol)) {
			t.Errorf("Three point out of tolerance.")
		}
	}
}

func BenchmarkTestPoint(b *testing.B) {
	f := testFunc
	for i := 0; i < b.N; i++ {
		ThreePoint(f, diff_test[4].x, h)
	}
}
