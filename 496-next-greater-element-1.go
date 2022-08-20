package main

// vanilla
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var ans = make([]int, len(nums1))
    for i, v := range nums1 {
        b := false
        ans[i] = -1
        for _, u := range nums2 {
            if b && u > v {
                ans[i] = u
                break
            }
            if v == u { b = true}
        }
    }
	return ans
}

// so-called monotone-stack
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
    var ans = make([]int, len(nums1))

    var next = make(map[int]int, len(nums2))
    var stack = make([]int, len(nums2))
    sp := 0
    for i := len(nums2)-1; i >=0; i-- {
        v := nums2[i]
        for sp > 0 && v > stack[sp-1] {
            sp--
        }

        if sp == 0 {
            next[v] = -1
        } else if v < stack[sp-1] {
            next[v] = stack[sp-1]
        }
        stack[sp] = v
        sp++
    }

    for i, v := range nums1 {
        ans[i] = next[v]
    }
	return ans
}
