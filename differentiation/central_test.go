package differentiation

import (
	"github.com/Arrow/nm/util"
	"math"
	"testing"
)

func TestCentral1(t *testing.T) {
	test_cases := diff_test[0].XYs()
	algo := NewCentral(diff_test[0].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 1: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentra2(t *testing.T) {
	test_cases := diff_test[1].XYs()
	algo := NewCentral(diff_test[1].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 2: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral3(t *testing.T) {
	test_cases := diff_test[2].XYs()
	algo := NewCentral(diff_test[2].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 3: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral4(t *testing.T) {
	test_cases := diff_test[3].XYs()
	algo := NewCentral(diff_test[3].f)
	algo.StepSize = 0.0000001
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 4: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral5(t *testing.T) {
	test_cases := diff_test[4].XYs()
	algo := NewCentral(diff_test[4].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 5: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral6(t *testing.T) {
	test_cases := diff_test[5].XYs()
	algo := NewCentral(diff_test[5].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 6: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral7(t *testing.T) {
	test_cases := diff_test[6].XYs()
	algo := NewCentral(diff_test[6].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 7: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral8(t *testing.T) {
	test_cases := diff_test[7].XYs()
	algo := NewCentral(diff_test[7].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 8: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral9(t *testing.T) {
	test_cases := diff_test[8].XYs()
	algo := NewCentral(diff_test[8].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 9: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral10(t *testing.T) {
	test_cases := diff_test[9].XYs()
	algo := NewCentral(diff_test[9].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 10: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral11(t *testing.T) {
	test_cases := diff_test[10].XYs()
	algo := NewCentral(diff_test[10].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 11: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral12(t *testing.T) {
	test_cases := diff_test[11].XYs()
	algo := NewCentral(diff_test[11].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 12: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func TestCentral13(t *testing.T) {
	test_cases := diff_test[12].XYs()
	algo := NewCentral(diff_test[12].f)
	for i, c := range test_cases {
		if !util.Tolerance(algo.Compute(c.X), c.Y, tol) {
			t.Errorf("Central out of tolerance. Function 13: i = %v, diff = %.6f, error = %.6f",
				i, math.Abs(c.Y-algo.Compute(c.X)), algo.er)
		}
	}
}

func BenchmarkCentral(b *testing.B) {
	algo := NewCentral(diff_test[0].f)
	for i := 0; i < b.N; i++ {
		algo.Compute(diff_test[0].a + 0.5*(diff_test[0].b-diff_test[0].a))
	}
}
