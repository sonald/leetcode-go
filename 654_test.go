package main

import (
	"fmt"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	var testcases = []struct {
		nums   []int
		expect *TreeNode
	}{
		{
			[]int{3, 2, 1, 6, 0, 5},
			BuildTreeFromList([]int{6, 3, 5, -1, 2, 0, -1, -1, 1}),
		},
	}

	for _, tc := range testcases {
		res := constructMaximumBinaryTree(tc.nums)
		if !compareTree(res, tc.expect) {
			tc.expect.Dump()
			fmt.Printf("\n")
			res.Dump()
			t.Errorf("nums: %v, tree is wrong\n", tc.nums)
		}

		res2 := constructMaximumBinaryTree2(tc.nums)
		if !compareTree(res2, tc.expect) {
			tc.expect.Dump()
			fmt.Printf("\n")
			res2.Dump()
			t.Errorf("nums: %v, tree is wrong\n", tc.nums)
		}
	}
}
