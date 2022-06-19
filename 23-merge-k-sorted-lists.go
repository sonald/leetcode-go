package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

// time: O(KNlgKN), space: O(KN)
func mergeKLists(lists []*ListNode) *ListNode {
	var res, head *ListNode
	var h NodeHeap
	heap.Init(&h)

	for _, l := range lists {
		for l != nil {
			heap.Push(&h, l)
			l = l.Next
		}
	}

	for h.Len() > 0 {
		nd := heap.Pop(&h).(*ListNode)
        nd.Next = nil
		if res == nil {
			res = nd
			head = nd
		} else {
			res.Next = nd
			res = res.Next
		}
	}

	return head
}

// sol2

type NodeWrapper struct {
    nd *ListNode
    index int
}

type NodeHeap2 []NodeWrapper

func (h NodeHeap2) Len() int {
	return len(h)
}

func (h NodeHeap2) Less(i, j int) bool {
	return h[i].nd.Val < h[j].nd.Val
}

func (h NodeHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap2) Push(x interface{}) {
	*h = append(*h, x.(NodeWrapper))
}

func (h *NodeHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	//old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

// time: O(KNlgK) space: O(K)
func mergeKLists2(lists []*ListNode) *ListNode {
	var res, head *ListNode
	var h NodeHeap2
	heap.Init(&h)

    for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
            heap.Push(&h, NodeWrapper{nd: lists[i], index: i})
		}
	}

	for h.Len() > 0 {
		nw := heap.Pop(&h).(NodeWrapper)
        nd := nw.nd
        next := nd.Next
        if next != nil {
            heap.Push(&h, NodeWrapper{nd: next, index: nw.index})
        }
		if res == nil {
			res = nd
			head = nd
		} else {
			res.Next = nd
			res = res.Next
		}
	}

	return head
}
