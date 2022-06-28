package main

import (
	"fmt"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	var testcases = []struct {
		nums     []int
		k        int
		t        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{[]int{-1, 2, -5, -3, 9, -1, 5, 9}, 2, 3, true},
		{[]int{1}, 2, 3, false},
		{[]int{1, 2}, 2, 3, true},
		{[]int{-3, 3}, 2, 4, false},
		{[]int{0}, 0, 0, false},
		{[]int{1, 2, 2, 3, 4, 5}, 3, 0, true},
		{[]int{2147483647, 2147483647 - 2, -2147483647}, 2, 2147483647, true},
		{[]int{2147483647, -2147483647}, 2, 2147483647, false},
		{[]int{1, 2}, 0, 1, false},
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := containsNearbyAlmostDuplicate(v.nums, v.k, v.t)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d_2", i), func(t *testing.T) {
			res := containsNearbyAlmostDuplicate2(v.nums, v.k, v.t)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d_3", i), func(t *testing.T) {
			res := containsNearbyAlmostDuplicate3(v.nums, v.k, v.t)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}
}
