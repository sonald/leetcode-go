package ds

type IntervalQueryTree struct {
	// map key is the index of the segment in the segment tree
	tree, lazy map[int]bool
	//range
	lo, hi int
}

func NewIntervalQueryTree(lo, hi int) *IntervalQueryTree {
	return &IntervalQueryTree{
		tree: make(map[int]bool),
		lazy: make(map[int]bool),
		lo:   lo,
		hi:   hi,
	}
}

func (st *IntervalQueryTree) queryAny(start, end, lo, hi, idx int) bool {
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
	return st.queryAny(start, end, lo, m, idx*2+1) || st.queryAny(start, end, m+1, hi, idx*2+2)
}

func (st *IntervalQueryTree) Query(start, end int) bool {
	return st.queryAny(start, end, st.lo, st.hi, 0)
}

func (st *IntervalQueryTree) update(start, end, lo, hi, idx int) {
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

func (st *IntervalQueryTree) Update(start, end int) {
	st.update(start, end, st.lo, st.hi, 0)
}
