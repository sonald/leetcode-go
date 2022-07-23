package main

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	var testcases = []struct {
		points   [][]int
		expected int
	}{
		{
			[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			2,
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			4,
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			2,
		},
		{
			[][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}},
			2,
		},
	}

	for i, tc := range testcases {
		res := findMinArrowShots(tc.points)
		if res != tc.expected {
			t.Errorf("round%d: expected: %v, real: %v\n", i, tc.expected, res)
		}

		res = findMinArrowShots2(tc.points)
		if res != tc.expected {
			t.Errorf("algo2: round%d: expected: %v, real: %v\n", i, tc.expected, res)
		}
	}
}
