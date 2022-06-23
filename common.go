package main

import (
	"container/list"
	"fmt"
	"sort"
)

func compareSlicesNoSort(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func countSort(nums []int, radix int, output []int) {
	var counts [10]int

	for i := 0; i < len(nums); i++ {
		v := (nums[i] % radix) / (radix / 10)
		counts[v]++
	}

	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		v := (nums[i] % radix) / (radix / 10)
		output[counts[v]-1] = nums[i]
		counts[v]--
	}

	copy(nums, output)
}

func max(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}

	return res
}

func radixSort(nums []int) {
	var maxVal = max(nums...)
	var output = make([]int, len(nums))

	for e := 1; e <= maxVal; e *= 10 {
		var counts [10]int

		for _, v := range nums {
			v = (v / e) % 10
			counts[v]++
		}

		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		for i := len(nums) - 1; i >= 0; i-- {
			v := (nums[i] / e) % 10
			output[counts[v]-1] = nums[i]
			counts[v]--
		}

		copy(nums, output)
	}
}

// Tree routines
//--------------------------------------------------------------------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// a list for a complete Tree (-1 for nil)
func BuildTreeFromList(data []int) *TreeNode {
	var bfs func(int) *TreeNode

	var stack = list.New()

	bfs = func(i int) *TreeNode {
		if i >= len(data) {
			return nil
		}
		if data[i] > 0 {
			return nil
		} // as nil pointer

		np := &TreeNode{Val: data[i]}
		np.Left = bfs(i + 1)
		np.Right = bfs(i + 2)

		return np
	}

	var i int = 0
	head := &TreeNode{Val: data[i]}
	i++
	stack.PushBack(head)
	for i < len(data) {
		np := stack.Remove(stack.Front()).(*TreeNode)
		if data[i] > 0 {
			np.Left = &TreeNode{Val: data[i]}
			stack.PushBack(np.Left)
		}
		i++
		if i < len(data) && data[i] > 0 {
			np.Right = &TreeNode{Val: data[i]}
			stack.PushBack(np.Right)
		}
		i++
	}

	return head
}

func (np *TreeNode) Dump() {
	fmt.Print("(")
	if np.Left != nil {
		np.Left.Dump()
	}

	fmt.Printf(" %d ", np.Val)

	if np.Right != nil {
		np.Right.Dump()
	}

	fmt.Print(")")
}

// AVL
//--------------------------------------------------------------------------------

type AVL struct {
	head *AVLNode
}

type AVLNode struct {
	Data        int
	h           int
	Left, Right *AVLNode
}

// rotate np right and return new subtree root
func rotateRight(np *AVLNode) *AVLNode {
	if np.Left == nil {
		return np
	}

	l := np.Left
	np.Left = l.Right
	l.Right = np
	np.updateHeight()
	l.updateHeight()
	return l
}

func rotateLeft(np *AVLNode) *AVLNode {
	if np.Right == nil {
		return np
	}

	r := np.Right
	np.Right = r.Left
	r.Left = np
	np.updateHeight()
	r.updateHeight()
	return r
}

func (np *AVLNode) leftLeaning() bool {
	return np.slant() >= 0
}

func (np *AVLNode) slant() int {
	var l, r int
	if np.Left != nil {
		l = np.Left.h
	}
	if np.Right != nil {
		r = np.Right.h
	}
	return l - r
}

func (np *AVLNode) rightLeaning() bool {
	return np.slant() <= 0
}

func (np *AVLNode) updateHeight() {
	var l, r int
	if np.Left != nil {
		l = np.Left.h
	}
	if np.Right != nil {
		r = np.Right.h
	}
	np.h = max(l, r) + 1
}

func insertNode(root *AVLNode, np *AVLNode) *AVLNode {
	if root == nil {
		return np
	}

	if root.Data > np.Data {
		root.Left = insertNode(root.Left, np)
	} else {
		root.Right = insertNode(root.Right, np)
	}
	root.updateHeight()

	if root.slant() >= 2 {
		target := root.Left
		if target.leftLeaning() {
			root = rotateRight(root)
		} else {
			root.Left = rotateLeft(target)
			root = rotateRight(root)
		}

	} else if root.slant() <= -2 {
		target := root.Right
		if target.rightLeaning() {
			root = rotateLeft(root)
		} else {
			root.Right = rotateRight(target)
			root = rotateLeft(root)
		}
	}
	root.updateHeight()

	return root
}

func (np *AVLNode) Dump() {
	fmt.Print("(")
	if np.Left != nil {
		np.Left.Dump()
	}

	fmt.Printf(" %d(%d) ", np.Data, np.h)

	if np.Right != nil {
		np.Right.Dump()
	}

	fmt.Print(")")
}

func (avl *AVL) Insert(data int) {
	np := &AVLNode{Data: data, h: 1}
	avl.head = insertNode(avl.head, np)
}

func (avl *AVL) Dump() {
	if avl.head != nil {
		avl.head.Dump()
	}
	fmt.Println("")
}

func AVLInit() *AVL {
	return &AVL{}
}
