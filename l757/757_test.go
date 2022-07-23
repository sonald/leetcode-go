package l757

import (
	"fmt"
	"testing"
)

func TestIntersectionSizeTwo(t *testing.T) {
	var testcases = []struct {
		intervals [][]int
		expected  int
	}{
		{
			[][]int{{1, 4}, {2, 5}, {1, 3}, {3, 5}},
			3,
		},
		{
			[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			5,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("round%d", i), func(t *testing.T) {
			res := intersectionSizeTwo(tc.intervals)
			if res != tc.expected {
				t.Errorf("intervals: %v, expected: %v, real: %v\n", tc.intervals, tc.expected, res)
			}
		})
	}
}
