package main

import (
	"fmt"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T) {
	var testcases = []struct {
		arr      []int
		expected [][]int
	}{
		{
			[]int{4, 2, 1, 3},
			[][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			[]int{1, 3, 6, 10, 15},
			[][]int{{1, 3}},
		},
		{
			[]int{3, 8, -10, 23, 19, -4, -14, 27},
			[][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := minimumAbsDifference2(tc.arr)
			if len(res) != len(tc.expected) {
				t.Fatalf("expected: %v, real: %v\n", tc.expected, res)
			}

			for i, v := range tc.expected {
				if !compareSlicesNoSort(v, res[i]) {
					t.Errorf("expected: %v, real: %v\n", tc.expected, res)
				}
			}
		})
	}
}
