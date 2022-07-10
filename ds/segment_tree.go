package ds

type SegmentTree struct {
	// map key is the index of the segment in the segment tree
	tree, lazy map[int]bool
	//range
	lo, hi int
}

func NewSegmentTree(lo, hi int) *SegmentTree {
	return &SegmentTree{
		tree: make(map[int]bool),
		lazy: make(map[int]bool),
		lo:   lo,
		hi:   hi,
	}
}

func (st *SegmentTree) query(start, end, lo, hi, idx int) bool {
	if start > hi || end < lo {
		return false
	}

	if st.lazy[idx] {
		return true
	}

	if start <= lo && end >= hi {
		return st.tree[idx]
	}

	m := lo + (hi-lo)/2
	return st.query(start, end, lo, m, idx*2+1) || st.query(start, end, m+1, hi, idx*2+2)
}

func (st *SegmentTree) Query(start, end int) bool {
	return st.query(start, end, st.lo, st.hi, 0)
}

func (st *SegmentTree) update(start, end, lo, hi, idx int) {
	if start > hi || end < lo {
		return
	}

	if start <= lo && end >= hi {
		st.tree[idx] = true
		st.lazy[idx] = true
	} else {
		m := lo + (hi-lo)/2
		st.update(start, end, lo, m, idx*2+1)
		st.update(start, end, m+1, hi, idx*2+2)
		st.tree[idx] = true
		if st.lazy[idx*2+1] && st.lazy[idx*2+2] {
			st.lazy[idx] = true
		}
	}
}

func (st *SegmentTree) Update(start, end int) {
	st.update(start, end, st.lo, st.hi, 0)
}
