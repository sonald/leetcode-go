package main

import (
	"fmt"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	var testcases = []struct {
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			1,
			[][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			4,
			[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			4 + 16*4,
			[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			9,
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			[][]int{{1}},
			100,
			[][]int{{1}},
		},
		{
			[][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}},
			23,
			[][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("round%d", i), func(t *testing.T) {
			res := shiftGrid2(tc.grid, tc.k)
			if !compareGrids(res, tc.expected) {
				t.Errorf("expected: %v, real: %v\n", tc.expected, res)
			}
		})
	}
}
