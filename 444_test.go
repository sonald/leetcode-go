package main

import "testing"

func TestSequenceReconstruction(t *testing.T) {
	var testcases = []struct {
		nums      []int
		sequences [][]int
		expected  bool
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2}, {1, 3}},
			false,
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2}},
			false,
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2}, {1, 3}, {2, 3}},
			true,
		},
		{
			[]int{1, 2, 4, 3},
			[][]int{{1, 2, 4}, {1, 3}},
			false,
		},
	}

	for i, tc := range testcases {
		res := sequenceReconstruction(tc.nums, tc.sequences)
		if res != tc.expected {
			t.Errorf("algo1: round%d: expected: %v, real: %v\n", i, tc.expected, res)
		}

		res = sequenceReconstruction2(tc.nums, tc.sequences)
		if res != tc.expected {
			t.Errorf("algo2: round%d: expected: %v, real: %v\n", i, tc.expected, res)
		}
	}
}
