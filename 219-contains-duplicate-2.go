package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
    var d = map[int][]int {}
    for i, v := range nums {
        d[v] = append(d[v], i)
	}

    for _, l := range d {
        if l == nil { continue}

        a := l[0]
        for _, b := range l[1:] {
            if b - a <= k {
                return true
            }
            a = b
        }
    }

	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
    var lastPos = make(map[int]int, len(nums)) // give capacity improves performance

    for i, v := range nums {
        if last, ok := lastPos[v]; ok && i - last <= k {
            return true
        }
        lastPos[v] = i
	}


	return false
}
