package main

import (
	"container/heap"
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

func rebalance(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}
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

func (np *AVLNode) getMin() *AVLNode {
	if np == nil {
		return nil
	}

	l := np
	for l.Left != nil {
		l = l.Left
	}
	return l
}

func deleteMinNode(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		root = root.Right
	} else {
		root.Left = deleteMinNode(root.Left)
	}

	if root != nil {
		root.updateHeight()
	}

	return rebalance(root)
}

func removeNode(root *AVLNode, data int) *AVLNode {
	if root == nil {
		return nil
	}

	if root.Data > data {
		root.Left = removeNode(root.Left, data)
	} else if root.Data < data {
		root.Right = removeNode(root.Right, data)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}

		md := root.Right.getMin()
		root.Right = deleteMinNode(root.Right)
		root.Data = md.Data
	}
	root.updateHeight()
	return rebalance(root)
}

func insertNodeIter(root *AVLNode, np *AVLNode) *AVLNode {
	if root == nil {
		return np
	}

	var sp = []**AVLNode{&root}
	for {
		p := sp[len(sp)-1]
		if (*p).Data > np.Data {
			if (*p).Left == nil {
				(*p).Left = np
				break
			}
			p = &(*p).Left
			sp = append(sp, p)
		} else {
			if (*p).Right == nil {
				(*p).Right = np
				break
			}
			p = &(*p).Right
			sp = append(sp, p)
		}
	}

	for len(sp) > 0 {
		p := sp[len(sp)-1]
		sp = sp[:len(sp)-1]

		(*p).updateHeight()
		*p = rebalance(*p)
	}

	return root
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

	return rebalance(root)
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

func (avl *AVL) getMin() int {
	return avl.head.getMin().Data
}

func (avl *AVL) Insert(data int) {
	np := &AVLNode{Data: data, h: 1}
	avl.head = insertNode(avl.head, np)
}

func (avl *AVL) Remove(data int) {
	avl.head = removeNode(avl.head, data)
}

func (avl *AVL) DeleteMin() {
	avl.head = deleteMinNode(avl.head)
}

func (avl *AVL) Dump() {
	if avl.head != nil {
		avl.head.Dump()
	}
	fmt.Println("")
}

func (avl *AVL) Size() int {
	var dfs func(node *AVLNode) int
	dfs = func(node *AVLNode) int {
		if node == nil {
			return 0
		}
		return dfs(node.Left) + dfs(node.Right) + 1
	}

	return dfs(avl.head)
}

func (avl *AVL) SearchDFS(pred func(v int) bool) *AVLNode {
	var dfs func(node *AVLNode) *AVLNode
	dfs = func(node *AVLNode) *AVLNode {
		if node == nil {
			return nil
		}

		if pred(node.Data) {
			lo := dfs(node.Left)
			if lo == nil {
				return node
			}

			return lo
		} else {
			return dfs(node.Right)
		}
	}

	return dfs(avl.head)
}

func (avl *AVL) Search(pred func(v int) bool) *AVLNode {
	if avl.head == nil {
		return nil
	}
	var stack = make([]*AVLNode, 0, avl.head.h)
	stack = append(stack, avl.head)
	for {
		np := stack[len(stack)-1]

		if pred(np.Data) {
			if np.Left == nil {
				return np
			}
			stack = append(stack, np.Left)
		} else {
			stack = stack[:len(stack)-1]

			if np.Right == nil {
				break
			}
			stack = append(stack, np.Right)
		}
	}

	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return nil
}

func AVLInit() *AVL {
	return &AVL{}
}

/// heap

type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Less(i, j int) bool { return a[i] > a[j] }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a *IntHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *IntHeap) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}

func HeapSort(a []int) {
	if len(a) <= 1 {
		return
	}

	h := IntHeap(a)
	heap.Init(&h)
	a[0], a[len(a)-1] = a[len(a)-1], a[0]

	for r := len(a) - 1; r > 0; r-- {
		h := IntHeap(a[0:r])
		heap.Fix(&h, 0)
		a[0], a[r-1] = a[r-1], a[0]
	}
}

///

func BucketSort(nums []int) {

}

/// Treap
type TreapNode struct {
	Left, Right *TreeNode
	prio        int
	Data        interface{}
}

type Treap struct {
	head *TreapNode
}
