package main

import (
	"math"
)

func min(a ...int) int {
	var res = math.MaxInt32
	for _, v := range a {
		if res > v {
			res = v
		}
	}

	return res
}

func paintHouseMinCost(costs [][]int) int {
	if costs == nil || len(costs) == 0 {
		return 0
	}

	var N = len(costs)
	var dp [3]int
	copy(dp[:], costs[0][:])

	for i := 1; i < N; i++ {
		var cur = [3]int{
			min(dp[1], dp[2]) + costs[i][0],
			min(dp[0], dp[2]) + costs[i][1],
			min(dp[0], dp[1]) + costs[i][2],
		}
		copy(dp[:], cur[:])
	}

	return min(dp[:]...)
}
