package main

import "testing"

func TestArrayNesting(t *testing.T) {
	t.Run("round1", func(t *testing.T) {
		var testcases = []struct {
			nums     []int
			expected int
		}{
			{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
			{[]int{0, 1, 2, 3, 4}, 1},
			{[]int{1, 2, 3, 4, 0}, 5},
		}

		for _, tc := range testcases {
			res := arrayNesting(tc.nums)
			if res != tc.expected {
				t.Errorf("expected: %v, real: %v\n", tc.expected, res)
			}
		}
	})

	t.Run("round2", func(t *testing.T) {
		var testcases = []struct {
			nums     []int
			expected int
		}{
			{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
			{[]int{0, 1, 2, 3, 4}, 1},
			{[]int{1, 2, 3, 4, 0}, 5},
		}

		for _, tc := range testcases {
			res := arrayNesting2(tc.nums)
			if res != tc.expected {
				t.Errorf("expected: %v, real: %v\n", tc.expected, res)
			}
		}
	})
}
