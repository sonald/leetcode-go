package main

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	N := len(nums)
	g := make([][]int, N+1)
	indeg := make([]int, N+1)
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			u, v := seq[i-1], seq[i]
			g[u] = append(g[u], v)
			indeg[v]++
		}
	}

	var q []int
	for i := 1; i <= N; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}

		u := q[0]
		q = nil
		for _, v := range g[u] {
			if indeg[v]--; indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return true
}

func sequenceReconstruction2(nums []int, sequences [][]int) bool {
	g := make(map[int]map[int]bool, len(nums))
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			u, v := seq[i-1], seq[i]
			if g[u] == nil {
				g[u] = make(map[int]bool)
			}
			g[u][v] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		u, v := nums[i-1], nums[i]
		if g[u] == nil || g[u][v] == false {
			return false
		}
	}

	return true
}
