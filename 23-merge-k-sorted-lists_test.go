package main

import (
	"testing"
)

func buildList(a []int) *ListNode {
	var head, res *ListNode

	for _, v := range a {
		nd := &ListNode{Val: v}
		if res == nil {
			head = nd
			res = nd
		} else {
			res.Next = nd
			res = res.Next
		}
	}

	return head
}

func compareList(l *ListNode, a []int, t *testing.T) {
	for i, v := range a {
		if l == nil || l.Val != v {
			t.Fatalf("#%d: expect %d real %d\n", i, v, l.Val)
		}
		l = l.Next
	}
}

func TestMergeKLists(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		lists := []*ListNode{
			buildList([]int{1, 4, 5}),
			buildList([]int{1, 3, 4}),
			buildList([]int{2, 6}),
		}

		res := mergeKLists(lists)
		if res == nil {
			t.Fatal("res == nil")
		}

		compareList(res, []int{1, 1, 2, 3, 4, 4, 5, 6}, t)
	})

	t.Run("two", func(t *testing.T) {
		lists := []*ListNode{}

		res := mergeKLists(lists)
		if res != nil {
			t.Fatal("res == nil")
		}
	})

	t.Run("three", func(t *testing.T) {
		var lists []*ListNode

		res := mergeKLists(lists)
		if res != nil {
			t.Fatal("res == nil")
		}
	})

	t.Run("4", func(t *testing.T) {
		lists := []*ListNode{
			buildList([]int{-2, -1, -1, -1}),
			nil,
		}

		res := mergeKLists(lists)
		if res == nil {
			t.Fatal("res == nil")
		}

		compareList(res, []int{-2, -1, -1, -1}, t)
	})
}

func TestMergeKLists2(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		lists := []*ListNode{
			buildList([]int{1, 4, 5}),
			buildList([]int{1, 3, 4}),
			buildList([]int{2, 6}),
		}

		res := mergeKLists2(lists)
		if res == nil {
			t.Fatal("res == nil")
		}

		compareList(res, []int{1, 1, 2, 3, 4, 4, 5, 6}, t)
	})

	t.Run("two", func(t *testing.T) {
		lists := []*ListNode{}

		res := mergeKLists2(lists)
		if res != nil {
			t.Fatal("res == nil")
		}
	})

	t.Run("three", func(t *testing.T) {
		var lists []*ListNode

		res := mergeKLists2(lists)
		if res != nil {
			t.Fatal("res == nil")
		}
	})

	t.Run("4", func(t *testing.T) {
		lists := []*ListNode{
			buildList([]int{-2, -1, -1, -1}),
			nil,
		}

		res := mergeKLists2(lists)
		if res == nil {
			t.Fatal("res == nil")
		}

		compareList(res, []int{-2, -1, -1, -1}, t)
	})
}
