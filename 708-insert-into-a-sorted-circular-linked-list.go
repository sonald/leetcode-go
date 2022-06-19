package main

type Node struct {
	Val  int
	Next *Node
}

// ref: leetcode 708
// O(N)
func insert(nd *Node, x int) *Node {
	res := &Node{Val: x}
	res.Next = res

	if nd == nil {
		return res
	}

	head := nd
	if nd.Next == nd {
	} else {
		for nd.Next != head {
			if x < nd.Val {
				if nd.Next.Val >= x && nd.Next.Val < nd.Val {
					break
				}

			} else {
				if nd.Next.Val >= x || nd.Val > nd.Next.Val {
					break
				}
			}

			nd = nd.Next
		}
	}

	res.Next = nd.Next
	nd.Next = res
	return head
}
