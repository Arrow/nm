package interval

import (
	"math"
)

type Interval struct {
	A, B float64
}

// Sum of x and y 
func (i *Interval) Add(x, y Interval) {
	i.A = x.A + y.A 
	i.B = x.B + y.B
}

// Subtraction of y from x
func (i *Interval) Sub(x, y Interval) {
	i.A = x.A - y.B
	i.B = x.B - y.A
}

// Multiplication of x and y
func (i *Interval) Mul(x, y Interval) {
	i.A = min([]float64{x.A*y.A, x.A*y.B, x.B*y.A, x.B*y.B})
	i.B = max([]float64{x.A*y.A, x.A*y.B, x.B*y.A, x.B*y.B})
}

// Division of x by y
// Special Case: If 0 is in range of y, returns NaN for both i.A and i.B
func (i *Interval) Div(x, y Interval) {
	if y.In(0.0) {
		i.A, i.B = math.NaN(), math.NaN()
		return
	}
	i.A = min([]float64{x.A/y.A, x.A/y.B, x.B/y.A, x.B/y.B})
	i.B = max([]float64{x.A/y.A, x.A/y.B, x.B/y.A, x.B/y.B})
}

// Returns true if x is within the interval
func (i Interval) In(x float64) bool {
	if x < i.A || x > i.B {
		return false
	}
	return true
}

func NewInterval(v float64) Interval {
	var i Interval
	i.A = math.Nextafter(v, math.Inf(-1))
	i.B = math.Nextafter(v, math.Inf(1))
	return i
}

func NewIntervalRange(a, b float64) Interval {
	var i Interval
	i.A, i.B = a, b
	return i
}

func min(fs []float64) float64 {
	var m float64
	for i, f := range fs {
		if i == 0 {
			m = f
		}
		if m > f {
			m = f
		}
	}
	return m
}

func max(fs []float64) float64 {
	var m float64
	for i, f := range fs {
		if i == 0 {
			m = f
		}
		if m < f {
			m = f
		}
	}
	return m
}
