package l715

import "fmt"

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

func (st *IntervalQueryTree) queryAll(start, end, lo, hi, idx int) bool {
	fmt.Printf("query: [%v, %v] [%v, %v] idx=%v tree=%v lazy=%v\n", start, end, lo, hi, idx, st.tree[idx], st.lazy[idx])

	//if start > hi || end < lo {
	//return true
	//}

	if start <= lo && end >= hi {
		if !st.lazy[idx] {
			fmt.Printf("query(false): [%v, %v] [%v, %v] idx=%v tree=%v lazy=%v\n", start, end, lo, hi, idx, st.tree[idx], st.lazy[idx])
			return false
		}
		return true
		//return st.lazy[idx]
	}

	m := lo + (hi-lo)/2
	if st.lazy[idx] {
		st.tree[idx*2+1] = true
		st.tree[idx*2+2] = true

		st.lazy[idx*2+1] = true
		st.lazy[idx*2+2] = true
		fmt.Printf("query: augment subs [%v, %v] [%v, %v] idx=%v tree=%v lazy=%v\n", start, end, lo, hi, idx, st.tree[idx], st.lazy[idx])
	}
	//return st.queryAll(start, end, lo, m, idx*2+1) && st.queryAll(start, end, m+1, hi, idx*2+2)

	if start <= m && !st.queryAll(start, end, lo, m, idx*2+1) {
		return false
	}

	if end > m && !st.queryAll(start, end, m+1, hi, idx*2+2) {
		return false
	}

	return true
}

func (st *IntervalQueryTree) Query(start, end int) bool {
	return st.queryAll(start, end, st.lo, st.hi, 0)
}

func (st *IntervalQueryTree) update(start, end, lo, hi, idx int) {
	if start > hi || end < lo {
		return
	}

	if start <= lo && end >= hi {
		st.tree[idx] = true
		st.lazy[idx] = true
		fmt.Printf("add1: [%v, %v] [%v, %v] idx=%v\n", start, end, lo, hi, idx)
	} else {
		m := lo + (hi-lo)/2
		st.update(start, end, lo, m, idx*2+1)
		st.update(start, end, m+1, hi, idx*2+2)
		st.tree[idx] = true
		if st.lazy[idx*2+1] && st.lazy[idx*2+2] {
			st.lazy[idx] = true
			fmt.Printf("add: [%v, %v] [%v, %v] idx=%v\n", start, end, lo, hi, idx)
		}
	}
}

func (st *IntervalQueryTree) Update(start, end int) {
	st.update(start, end, st.lo, st.hi, 0)
}

func (st *IntervalQueryTree) remove(start, end, lo, hi, idx int) {
	if start > hi || end < lo {
		return
	}

	fmt.Printf("before remove: [%v, %v] [%v, %v] idx=%v, tree=%v, lazy=%v\n", start, end, lo, hi, idx,
		st.tree[idx], st.lazy[idx])

	if start <= lo && end >= hi {
		st.tree[idx] = false
		st.lazy[idx] = false

		fmt.Printf("remove1: [%v, %v] [%v, %v] idx=%v \n", start, end, lo, hi, idx)
	} else {
		m := lo + (hi-lo)/2
		if st.lazy[idx] {
			st.tree[idx*2+1] = true
			st.tree[idx*2+2] = true

			st.lazy[idx*2+1] = true
			st.lazy[idx*2+2] = true
			fmt.Printf("mark true: [%v, %v] [%v, %v] idx=%v, lazy=%v\n", start, end, lo, hi, idx, st.lazy[idx])
		}

		st.remove(start, end, lo, m, idx*2+1)
		st.remove(start, end, m+1, hi, idx*2+2)

		//if !st.lazy[idx*2+1] && !st.lazy[idx*2+2] {
		if !st.tree[idx*2+1] && !st.tree[idx*2+2] {
			//fmt.Printf("remove2: [%v, %v] [%v, %v] idx=%v, tree=%v, lazy=%v\n", start, end, lo, hi, idx,
			//st.tree[idx], st.lazy[idx])
			st.tree[idx] = false
		}
		if !st.lazy[idx*2+1] || !st.lazy[idx*2+2] {
			st.lazy[idx] = false
		}
		fmt.Printf("after remove: [%v, %v] [%v, %v] idx=%v, tree=%v, lazy=%v\n", start, end, lo, hi, idx,
			st.tree[idx], st.lazy[idx])
	}
}

func (st *IntervalQueryTree) Remove(start, end int) {
	st.remove(start, end, st.lo, st.hi, 0)
}

type RangeModule struct {
	tree *IntervalQueryTree
}

func Constructor() RangeModule {
	return RangeModule{
		tree: NewIntervalQueryTree(1, 1e9),
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.tree.Update(left, right-1)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.tree.Query(left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.tree.Remove(left, right-1)
}
