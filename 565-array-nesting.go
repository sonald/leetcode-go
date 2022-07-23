package main

// mark all loops distinctly
func arrayNesting(nums []int) int {
	res := 0
	for i, k := range nums {
		count := 1
		if i != k && nums[k] != k {
			next := k
			nums[i] = i
			for next != nums[next] {
				nums[next], next = i, nums[next]
				count++
			}
		}
		if count > res {
			res = count
		}
	}
	return res
}

func arrayNesting2(nums []int) int {
	res, N := 0, len(nums)
	for i := range nums {
		count := 0
		for nums[i] < N {
			nums[i], i = N, nums[i]
			count++
		}
		if count > res {
			res = count
		}
	}

	return res
}
