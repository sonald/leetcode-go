package l346

import (
	"math"
	"testing"
)

func TestMovingAverage_Next(t *testing.T) {
	c := Constructor(3)
	var testcases = []struct {
		next     int
		expected float64
	}{
		{1, 1.0},
		{10, 5.5},
		{3, 4.66667},
		{5, 6.0},
	}
	for _, tc := range testcases {
		res := c.Next(tc.next)
		if math.Abs(tc.expected-res) > 0.00001 {
			t.Errorf("expected: %v, real: %v\n", tc.expected, res)
		}
	}
}
