package differentiation

import (
	"math"
	"testing"
	"github.com/Arrow/nm/util"
)

var diff_test3pt = []util.TestXY{
	{-5, -10},
	{-2, -4},
	{0, 0},
	{2, 4},
	{5, 10},
}

func testFunc3pt(x float64) float64 {
	return math.Pow(x, 2)
}

func TestThreePoint(t *testing.T) {
	f := testFunc3pt
	for _, c := range diff_test3pt {
		if(!util.Tolerance(ThreePoint(f, c.X), c.Y, tol)) {
			t.Errorf("Three point out of tolerance.")
		}
	}
}

func BenchmarkThreePoint(b *testing.B) {
	f := testFunc3pt
	for i := 0; i < b.N; i++ {
		ThreePoint(f, diff_test3pt[4].X)
	}
}
