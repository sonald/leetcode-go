package l715

import (
	"fmt"
	"github.com/sonald/leetcode-go/ds"
)

type Interval struct {
	start, end int
}

func (it Interval) Compare(other ds.AVLInterface) int {
	it2 := other.(Interval)
	if it.start < it2.start {
		return -1
	} else if it.start > it2.start {
		return 1
	}

	if it.end < it2.end {
		return -1
	} else if it.end == it2.end {
		return 0
	} else {
		return 1
	}
}

type RangeModule2 struct {
	intervals *ds.AVL
}

func Constructor2() RangeModule2 {
	return RangeModule2{
		intervals: ds.AVLInit(),
	}
}

func (this *RangeModule2) AddRange(left int, right int) {
	target := Interval{start: left, end: right - 1}
	np := this.intervals.RevSearch(func(v ds.AVLInterface) bool {
		return target.start >= v.(Interval).start
	})

	fmt.Printf("add(%v): np=%v\n", target, np)
	if np != nil {
		it := np.Data.(Interval)
		if target.end <= it.end {
			return
		}

		if it.end >= target.start-1 {
			target.start = it.start

			data := np.Data
			np = this.intervals.Successor(np)
			this.intervals.Remove(data)
		} else {
			np = this.intervals.Successor(np)
		}
	} else {
		np = this.intervals.GetMinNode()
	}

	fmt.Printf("add2(%v): np=%v\n", target, np)
	for np != nil {
		it := np.Data.(Interval)
		if it.start > target.end+1 {
			break
		}

		if target.end < it.end {
			target.end = it.end
		}

		data := np.Data
		np = this.intervals.Successor(np)
		this.intervals.Remove(data)
	}

	this.intervals.Insert(target)

	this.intervals.Dump()
}

func (this *RangeModule2) QueryRange(left int, right int) bool {
	it := Interval{start: left, end: right - 1}
	np := this.intervals.RevSearch(func(v ds.AVLInterface) bool {
		return it.start >= v.(Interval).start
	})
	fmt.Printf("q(%v) np = %v\n", it, np)
	this.intervals.Dump()
	return np != nil && it.end <= np.Data.(Interval).end
}

func (this *RangeModule2) RemoveRange(left int, right int) {
	target := Interval{start: left, end: right - 1}
	np := this.intervals.RevSearch(func(v ds.AVLInterface) bool {
		return target.start >= v.(Interval).start
	})

	if np == nil {
		np = this.intervals.GetMinNode()
	} else {
		it := np.Data.(Interval)
		if it.end < target.start {
			np = this.intervals.Successor(np)
		} else if it.end >= target.end {
			it2 := it
			it2.start = target.end + 1
			if it2.start <= it2.end {
				this.intervals.Insert(it2)
			}

			it.end = target.start - 1
			if it.start <= it.end {
				np.Data = it
			} else {
				this.intervals.Remove(np.Data)
			}
			return
		} else {
			it.end = target.start - 1
			if it.start <= it.end {
				np.Data = it
				np = this.intervals.Successor(np)
			} else {
				data := np.Data
				np = this.intervals.Successor(np)
				this.intervals.Remove(data)
			}
		}

	}

	fmt.Printf("rm(%v): np =%v\n", target, np)
	for np != nil {
		it := np.Data.(Interval)
		if it.start > target.end {
			break
		}

		if it.end <= target.end {
			data := np.Data
			np = this.intervals.Successor(np)
			this.intervals.Remove(data)
			//next
		} else {
			it.start = target.end + 1
			np.Data = it
			break
		}
	}

	this.intervals.Dump()
}
