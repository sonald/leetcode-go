package main

import (
	"fmt"
	"testing"
)

func TestPaintHouseMinCost(t *testing.T) {
	var testcases = []struct {
		costs    [][]int
		expected int
	}{
		{
			[][]int{
				{17, 2, 17},
				{16, 16, 5},
				{14, 3, 19},
			}, 10,
		},
		{
			[][]int{
				{7, 6, 2},
			}, 2,
		},
		{
			nil,
			0,
		},
		{
			[][]int{
				{11, 1, 17},
				{16, 16, 5},
				{14, 3, 19},
				{12, 1, 14},
			}, 21,
		},
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := paintHouseMinCost(v.costs)
			if res != v.expected {
				t.Errorf("expected: %v, real: %v\n", v.expected, res)
			}
		})
	}
}
