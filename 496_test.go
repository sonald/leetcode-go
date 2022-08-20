package main

import (
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	var testcases = []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}, []int{7, 7, 7, 7, 7}},
	}

	for _, tc := range testcases {
		res := nextGreaterElement2(tc.nums1, tc.nums2)
		if !compareSlicesNoSort(res, tc.expect) {
			t.Errorf("expected: %v, real: %v\n", tc.expect, res)
		}
	}
}
