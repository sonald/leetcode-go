package main

import (
	"sort"
)

func indexOf(a []int, v2 int) int {
	for i, v := range a {
		if v == v2 {
			return i
		}
	}

	return -1
}

//time: O(NlgN), space: O(N)
func twoSum(nums []int, target int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	sort.Ints(res)

	for i, v := range res {
		var half = target - v
		var l, m = i + 1, len(res) - 1
		var k = -1

	inner:
		for l <= m {
			k = l + (m-l)/2
			switch {
			case res[k] == half:
				break inner // found it
			case res[k] < half:
				l = k + 1
			default:
				m = k - 1
			}
		}

		if l <= m {
			id1 := indexOf(nums, res[i])
			nums[id1] += -1
			id2 := indexOf(nums, res[k])
			return []int{id1, id2}

		}
	}

	return nil
}

//time: O(N), space: O(N)
func twoSum2(nums []int, target int) []int {
	var m = map[int]int{}

	for i, v := range nums {
		if _, ok := m[v]; ok {
			return []int{m[v], i}
		}

		m[target-v] = i
	}

	return nil
}
