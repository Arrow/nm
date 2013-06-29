package differentiation

import (
	"github.com/Arrow/nm/util"
	"math"
	"testing"
)

func TestThreePoint(t *testing.T) {
	test_cases := diff_test[0].XYs()
	for i, c := range test_cases {
		if !util.Tolerance(ThreePoint(diff_test[0].f, c.X), c.Y, tol) {
			t.Errorf("Three point out of tolerance. i = %v, diff = %.6f",
				i, math.Abs(c.Y-c.X))
		}
	}
}

func BenchmarkThreePoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreePoint(diff_test[0].f, diff_test[0].b)
	}
}
