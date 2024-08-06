package b

import (
	"math"
	"testing"
)

var tcs = []struct {
	m         []float64
	varExpect float64
}{
	{[]float64{1, 2, 3, 4}, 1.25},
	{[]float64{610, 450, 160, 420, 310}, 22440},
}

func compare(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}

func TestVariance(t *testing.T) {
	for _, tc := range tcs {
		res := Variance(tc.m)
		expect := tc.varExpect

		if !compare(res, expect) {
			t.Errorf("Variance: %f != %f", res, expect)
		}
	}
}
