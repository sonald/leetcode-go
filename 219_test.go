package main

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	var testcases = []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := containsNearbyDuplicate(v.nums, v.k)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d_2", i), func(t *testing.T) {
			res := containsNearbyDuplicate2(v.nums, v.k)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d_2", i), func(t *testing.T) {
			res := containsNearbyDuplicate3(v.nums, v.k)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}
}
