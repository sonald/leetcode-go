package main

// O(NK)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, v := range nums {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			a, b := v, nums[j]
			if a > b {
				a, b = b, a
			}
			if b-a <= t {
				return true
			}
		}
	}
	return false
}

// O(NlgK)
// build a tree with k nodes and sliding the tree
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	avl := AVLInit()

	for i := 0; i < len(nums); i++ {
		np := avl.Search(func(v int) bool {
			return v >= nums[i]-t
		})

		if np != nil {
			if np.Data-nums[i] <= t {
				return true
			}
		}

		avl.Insert(nums[i])
		if i >= k {
			avl.Remove(nums[i-k])
		}
	}

	return false
}

func containsNearbyAlmostDuplicate3(nums []int, k int, t int) bool {
	var buckets = map[int]int{}

	var ID = func(x int) int {
		if x < 0 {
			return (x+1)/(t+1) - 1
		}
		return x / (t + 1)
	}

	for i := 0; i < len(nums); i++ {
		var id = ID(nums[i])
		if _, ok := buckets[id]; ok {
			return true
		}
		buckets[id] = nums[i]
		if v, ok := buckets[id-1]; ok && (nums[i]-v) <= t {
			return true
		}
		if v, ok := buckets[id+1]; ok && (v-nums[i]) <= t {
			return true
		}

		if i >= k {
			delete(buckets, ID(nums[i-k]))
		}
	}

	return false
}
