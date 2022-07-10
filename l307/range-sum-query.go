package l307

import "github.com/sonald/leetcode-go/ds"

type NumArray struct {
	st *ds.SegmentTree
}

func Constructor(nums []int) NumArray {
	st := ds.NewSegmentTree(nums)

	return NumArray{
		st: st,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.st.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.st.SumRangeIter(left, right)
}
