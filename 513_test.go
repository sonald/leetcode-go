package main

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
	var cases = []struct {
		data     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7}, 7},
		{[]int{2, 1, 3}, 1},
		{[]int{1}, 1},
	}

	for _, v := range cases {
		t.Run("one", func(t *testing.T) {
			root := BuildTreeFromList(v.data)
			//root.Dump()
			res := findBottomLeftValue3(root)
			if res != v.expected {
				t.Errorf("expected: %d, real: %d\n", v.expected, res)
			}
		})
	}
}
