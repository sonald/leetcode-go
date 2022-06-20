package main

import (
	"testing"
)

func TestFindFrequentTreeSum(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		root := &TreeNode{
			5,
			&TreeNode{2, nil, nil},
			&TreeNode{-3, nil, nil},
		}

		res := findFrequentTreeSum(root)
		expect := []int{2, -3, 4}
		if !compareSlices(res, expect) {
			t.Fatalf("expected: %v, real: %v\n", expect, res)
		}
	})

	t.Run("two", func(t *testing.T) {
		root := &TreeNode{
			5,
			&TreeNode{2, nil, nil},
			&TreeNode{-5, nil, nil},
		}

		res := findFrequentTreeSum(root)
		expect := []int{2}
		if !compareSlices(res, expect) {
			t.Fatalf("expected: %v, real: %v\n", expect, res)
		}
	})
}
