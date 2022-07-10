package l729

import (
	"github.com/sonald/leetcode-go/ds"
	"sort"
)

type Interval729 struct {
	start int
	end   int
}

type IntervalList729 []Interval729

type MyCalendar struct {
	intervals IntervalList729
	bst       *ds.AVL
	st        *ds.IntervalQueryTree
}

func Constructor() MyCalendar {
	return MyCalendar{
		bst: ds.AVLInit(),
		st:  ds.NewIntervalQueryTree(0, 1e9),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	interval := Interval729{start, end}
	j := sort.Search(len(this.intervals), func(i int) bool {
		return interval.start <= this.intervals[i].start
	})

	if j == len(this.intervals) {
		if j > 0 && interval.start < this.intervals[j-1].end {
			return false
		}
		this.intervals = append(this.intervals, interval)
	} else {
		if interval.end > this.intervals[j].start || (j > 0 && interval.start < this.intervals[j-1].end) {
			return false
		}
		newc := make([]Interval729, len(this.intervals)+1)
		copy(newc[:j], this.intervals[:j])
		newc[j] = interval
		copy(newc[j+1:], this.intervals[j:])
		this.intervals = newc

	}

	return true
}

func (it Interval729) Compare(i ds.AVLInterface) int {
	it2 := i.(Interval729)
	if it.start > it2.start {
		return 1
	} else if it.start < it2.start {
		return -1
	}
	return 0
}

func (this *MyCalendar) Book2(start int, end int) bool {
	interval := Interval729{start, end}

	np := this.bst.RevSearch(func(i ds.AVLInterface) bool {
		it := i.(Interval729)
		return interval.start > it.start
	})

	if np != nil {
		if interval.start < np.Data.(Interval729).end {
			return false
		}

		next := this.bst.Successor(np)
		if next != nil && interval.end > next.Data.(Interval729).start {
			return false
		}

	} else {
		if this.bst.Size() > 0 {
			it := this.bst.GetMin().(Interval729)
			if interval.end > it.start {
				return false
			}
		}
	}

	this.bst.Insert(interval)
	return true
}

func (this *MyCalendar) Book3(start int, end int) bool {
	if this.st.Query(start, end-1) {
		return false
	}

	this.st.Update(start, end-1)
	return true
}
