package ds

import (
	"fmt"
	"testing"
)

type Int int

func (i Int) Compare(j AVLInterface) int {
	if i > j.(Int) {
		return 1
	} else if i < j.(Int) {
		return -1
	}
	return 0
}

func TestAVL_Successor(t *testing.T) {
	var data = []int{3, 1, 10, 4, 8, 5, 6, 11, 9}
	avl := AVLInit()
	for _, v := range data {
		avl.Insert(Int(v))
	}
	avl.Dump()

	var testcases = []struct {
		bound Int
		curr  Int
		next  Int
	}{
		{2, 3, 4},
		{7, 8, 9},
		{1, 1, 3},
		{0, 1, 3},
		{5, 5, 6},
		{4, 4, 5},
		{6, 6, 8},
		{11, 11, -1},
	}

	for _, p := range testcases {
		np := avl.Search(func(v AVLInterface) bool {
			return v.(Int) >= p.bound
		})

		next := avl.Successor(np)
		if p.next == -1 {
			if next != nil {
				t.Errorf("NIL: bound: %d, expected: %d, real: %+v\n", p.bound, p.next, next)
			}
		} else if np == nil || next == nil || next.Data != p.next {
			t.Errorf("bound: %d, expected: %d, real: %+v\n", p.bound, p.next, next)
		}
	}
}

func TestAVL_RevSearch(t *testing.T) {
	var data = []int{3, 1, 10, 4, 8, 5, 6, 11, 9}
	avl := AVLInit()
	for _, v := range data {
		avl.Insert(Int(v))
	}
	avl.Dump()

	var testcases = []struct {
		bound    Int
		expected Int
	}{
		{2, 1},
		{7, 6},
		{1, 1},
		{0, -1},
		{12, 11},
		{-12, -1},
		{5, 5},
		{4, 4},
	}

	for _, p := range testcases {
		np := avl.RevSearch(func(v AVLInterface) bool {
			return v.(Int) <= p.bound
		})

		if p.expected < 0 {
			if np != nil {
				t.Errorf("expected: nil, real: %+v\n", np)
			}
		} else {
			if np == nil || np.Data != p.expected {
				t.Errorf("bound: %d, expected: %d, real: %+v\n", p.bound, p.expected, np)
			}
		}
	}
}

func TestNewIntervalQueryTree(t *testing.T) {
	type Interval struct {
		start, end int
	}
	t.Run("round1", func(t *testing.T) {
		st := NewIntervalQueryTree(0, 1e9)
		var testcases = []struct {
			interval Interval
			result   bool
		}{
			{Interval{10, 20}, false},
			{Interval{15, 25}, true},
			{Interval{20, 30}, false},
			{Interval{40, 50}, false},
			{Interval{30, 42}, true},
		}

		for _, tc := range testcases {
			res := st.Query(tc.interval.start, tc.interval.end-1)
			if res != tc.result {
				t.Fatalf("query: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
			}

			if !res {
				st.Update(tc.interval.start, tc.interval.end-1)
			}
		}
	})

	t.Run("round2", func(t *testing.T) {
		st := NewIntervalQueryTree(0, 1e9)
		var testcases = []struct {
			interval Interval
			result   bool
		}{
			{Interval{37, 50}, false},
			{Interval{33, 50}, true},
			{Interval{4, 17}, false},
			{Interval{35, 48}, true},
			{Interval{8, 25}, true},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := st.Query(tc.interval.start, tc.interval.end-1)
				if res != tc.result {
					t.Errorf("expected: %v, real: %v\n", tc.result, res)
				}
				if !res {
					st.Update(tc.interval.start, tc.interval.end-1)
				}
			})
		}
	})

	t.Run("round3", func(t *testing.T) {
		st := NewIntervalQueryTree(0, 1e9)
		var testcases = []struct {
			interval Interval
			result   bool
		}{
			{Interval{47, 50}, false},
			{Interval{33, 41}, false},
			{Interval{39, 45}, true},
			{Interval{33, 42}, true},
			{Interval{25, 32}, false},
			{Interval{26, 35}, true},
			{Interval{19, 25}, false},
			{Interval{3, 8}, false},
			{Interval{8, 13}, false},
			{Interval{18, 27}, true},
		}

		for _, tc := range testcases {
			res := st.Query(tc.interval.start, tc.interval.end-1)
			if res != tc.result {
				t.Fatalf("query: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
			}
			if !res {
				st.Update(tc.interval.start, tc.interval.end-1)
			}
		}
	})

}
