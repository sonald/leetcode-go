package l307

import (
	"fmt"
	"testing"
)

func TestNumArray(t *testing.T) {
	type Op struct {
		cmd  byte
		args []int
	}

	var testcases = []struct {
		nums []int
		op   []Op
	}{
		{
			[]int{1, 3, 5},
			[]Op{
				{'s', []int{0, 2, 9}},
				{'u', []int{1, 2}},
				{'s', []int{0, 2, 8}},
			},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]Op{
				{'s', []int{0, 11, 66}},
				{'s', []int{0, 7, 28}},
				{'s', []int{1, 7, 28}},
				{'s', []int{2, 7, 27}},
				{'s', []int{3, 8, 33}},
				{'s', []int{3, 8, 33}},
				{'s', []int{3, 10, 52}},
				{'u', []int{8, 9}},
				{'s', []int{3, 10, 53}},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			na := Constructor(tc.nums)
			for _, op := range tc.op {
				switch op.cmd {
				case 's':
					res := na.SumRange(op.args[0], op.args[1])
					if res != op.args[2] {
						t.Fatalf("sumRange: expected: %v, rea: %v\n", op.args[2], res)
					}

				case 'u':
					na.Update(op.args[0], op.args[1])
				}
			}
		})
	}
}
