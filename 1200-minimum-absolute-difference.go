package main

import (
	"math"
	"sort"
)

type IntPairs [][]int

func (p IntPairs) Len() int           { return len(p) }
func (p IntPairs) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p IntPairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//O(N^2)
func minimumAbsDifference(arr []int) [][]int {
	var res IntPairs
	m := math.MaxInt32

	for i := 0; i < len(arr)-1; i++ {
		for _, v2 := range arr[i+1:] {
			v1 := arr[i]
			if v1 > v2 {
				v1, v2 = v2, v1
			}

			if v2-v1 < m {
				m = v2 - v1
				res = [][]int{{v1, v2}}
			} else if v2-v1 == m {
				res = append(res, []int{v1, v2})
			}
		}
	}

	sort.Sort(res)
	return res
}

//O(NlgN)
func minimumAbsDifference2(arr []int) [][]int {
	var res IntPairs
	m := math.MaxInt32

	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		v2, v1 := arr[i], arr[i-1]
		if v2-v1 < m {
			m = v2 - v1
			res = [][]int{{v1, v2}}
		} else if v2-v1 == m {
			res = append(res, []int{v1, v2})
		}
	}

	sort.Sort(res)
	return res
}
