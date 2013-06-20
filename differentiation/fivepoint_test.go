package differentiation

import (
	"math"
	"testing"
	"github.com/Arrow/nm/util"
)

var diff_test5pt = []util.TestXY{
	{-5, -10},
	{-2, -4},
	{0, 0},
	{2, 4},
	{5, 10},
}

func testFunc5pt(x float64) float64 {
	return math.Pow(x, 2)
}

func TestFivePoint(t *testing.T) {
	f := testFunc5pt
	for _, c := range diff_test5pt {
		if(!util.Tolerance(FivePoint(f, c.X), c.Y, tol)) {
			t.Errorf("Three point out of tolerance.")
		}
	}
}

func BenchmarkFivePoint(b *testing.B) {
	f := testFunc5pt
	for i := 0; i < b.N; i++ {
		FivePoint(f, diff_test5pt[4].X)
	}
}
