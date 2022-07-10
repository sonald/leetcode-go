package ds

import (
	"math"
)

type SegmentTree struct {
	N    int
	data []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	x := math.Ceil(math.Log2(float64(len(nums))))
	x = math.Pow(2, x)
	n := int(x)

	st := &SegmentTree{
		N:    n,
		data: make([]int, 2*n),
	}

	for i := 0; i < len(nums); i++ {
		st.data[i+n] = nums[i]
	}

	for i := n - 1; i > 0; i-- {
		st.data[i] = st.data[i*2] + st.data[i*2+1]
	}

	return st
}

// i is index of original array
func (st *SegmentTree) Update(i, v int) {
	j := i + st.N
	st.data[j] = v
	for {
		j /= 2
		if j <= 0 {
			break
		}
		st.data[j] = st.data[j*2] + st.data[j*2+1]
	}
}

func (st *SegmentTree) sumRange(i, j, l, h, idx int) int {
	if i > h || j < l {
		return 0
	}
	//fmt.Printf("%d,%d  %d,%d  %d\n", i, j, l, h, idx)

	if i <= l && j >= h {
		return st.data[idx]
	}

	m := l + (h-l+1)/2
	return st.sumRange(i, j, l, m-1, idx*2) + st.sumRange(i, j, m, h, idx*2+1)
}

// i,j are indice of original array
func (st *SegmentTree) SumRange(i, j int) int {
	return st.sumRange(i, j, 0, st.N-1, 1)
}

// idea: https://codeforces.com/blog/entry/18051
// but query impl above seems wrong?
// and this https://codeforces.com/blog/entry/1256
func (st *SegmentTree) SumRangeIter(i, j int) int {
	res := 0
	for l, r := i+st.N, j+st.N; l <= r; l, r = l/2, r/2 {
		if (l & 1) == 1 {
			res += st.data[l]
			l++
		}

		if (r & 1) == 0 {
			res += st.data[r]
			r--
		}
	}
	return res
}
