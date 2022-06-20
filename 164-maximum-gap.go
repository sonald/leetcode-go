package main

import (
	"sort"
)

// time: O(NlgN) space: O(1)
func maximumGap(nums []int) int {
	sort.Ints(nums)
	var max = 0

	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		if gap > max {
			max = gap
		}
	}
	return max
}

func maximumGap2(nums []int) int {
	radixSort(nums)
	var res = 0
	for i := 1; i < len(nums); i++ {
		res = max(nums[i]-nums[i-1], res)
	}
	return res
}
