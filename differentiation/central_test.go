package differentiation

import (
	"github.com/Arrow/nm/util"
	"math"
	"testing"
)

func TestCentral(t *testing.T) {
	test_cases := diff_test[0].XYs()
	algo := NewCentral(diff_test[0].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Five point out of tolerance. i = %v, diff = %.6f",
				i, math.Abs(c.Y-c.X))
		}
	}
}

func BenchmarkCentral(b *testing.B) {
	algo := NewCentral(diff_test[0].f)
	for i := 0; i < b.N; i++ {
		algo.Compute(diff_test[0].b)
	}
}
