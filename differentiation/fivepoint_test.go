package differentiation

import (
	"github.com/Arrow/nm/util"
	"math"
	"testing"
)

func TestFivePoint(t *testing.T) {
	test_cases := diff_test[0].XYs()
	for i, c := range test_cases {
		if !util.Tolerance(FivePoint(diff_test[0].f, c.X), c.Y, tol) {
			t.Errorf("Five point out of tolerance. i = %v, diff = %.6f",
				i, math.Abs(c.Y-c.X))
		}
	}
}

func BenchmarkFivePoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FivePoint(diff_test[0].f, diff_test[0].b)
	}
}
