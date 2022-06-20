package main

import (
	"math"
	"sort"
)

func compareSortedSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func countSort(nums []int, radix int, output []int) {
	var counts [10]int

	for i := 0; i < len(nums); i++ {
		v := (nums[i] % radix) / (radix / 10)
		counts[v]++
	}

	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		v := (nums[i] % radix) / (radix / 10)
		output[counts[v]-1] = nums[i]
		counts[v]--
	}

	copy(nums, output)
}

func max(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}

	return res
}

func radixSort(nums []int) {
	var maxVal = max(nums...)
	var output = make([]int, len(nums))

	for e := 1; e <= maxVal; e *= 10 {
		var counts [10]int

		for _, v := range nums {
			v = (v / e) % 10
			counts[v]++
		}

		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		for i := len(nums) - 1; i >= 0; i-- {
			v := (nums[i] / e) % 10
			output[counts[v]-1] = nums[i]
			counts[v]--
		}

		copy(nums, output)
	}
}

func radixSort2(nums []int) {
	var max = 0
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	maxe := 1
	for max/int(math.Pow10(maxe)) > 0 {
		maxe++
	}

	var output = make([]int, len(nums))

	for e := 1; e <= maxe; e++ {
		countSort(nums, int(math.Pow10(e)), output)
		//log.Printf("count: %v\n", nums)
	}
}
