package ds

import (
	"fmt"
	"math"
)

func max(v ...int) int {
	res := math.MinInt32

	for _, d := range v {
		if res < d {
			res = d
		}
	}

	return res
}

type AVLInterface interface {
	Compare(v AVLInterface) int
}

type AVL struct {
	head *AVLNode
}

type AVLNode struct {
	Data        AVLInterface
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

func removeNode(root *AVLNode, data AVLInterface) *AVLNode {
	if root == nil {
		return nil
	}

	switch root.Data.Compare(data) {
	case 1:
		root.Left = removeNode(root.Left, data)
	case -1:
		root.Right = removeNode(root.Right, data)
	default:
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

func insertNode(root *AVLNode, np *AVLNode) *AVLNode {
	if root == nil {
		return np
	}

	switch root.Data.Compare(np.Data) {
	case 1:
		root.Left = insertNode(root.Left, np)
	default:
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

	fmt.Printf(" %v(%d) ", np.Data, np.h)

	if np.Right != nil {
		np.Right.Dump()
	}

	fmt.Print(")")
}

func (avl *AVL) GetMin() AVLInterface {
	return avl.head.getMin().Data
}

func (avl *AVL) Parent(np *AVLNode) *AVLNode {
	var par *AVLNode
	cur := avl.head
	for {
		if cur == np {
			return par
		}

		par = cur
		switch cur.Data.Compare(np.Data) {
		case 1:
			cur = cur.Left
		case -1:
			cur = cur.Right
		}
	}
}

func (avl *AVL) Successor(np *AVLNode) *AVLNode {
	var sp []*AVLNode
	if np.Right == nil {
		cur := avl.head
		for {
			if cur == np {
				for i := len(sp) - 1; i >= 0; i-- {
					if sp[i].Left == cur {
						return sp[i]
					}
					cur = sp[i]
				}

				return nil
			}

			sp = append(sp, cur)
			switch cur.Data.Compare(np.Data) {
			case 1:
				cur = cur.Left
			case -1:
				cur = cur.Right
			}
		}
	}
	return np.Right.getMin()
}

func (avl *AVL) Insert(data AVLInterface) {
	np := &AVLNode{Data: data, h: 1}
	avl.head = insertNode(avl.head, np)
}

func (avl *AVL) Remove(data AVLInterface) {
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

func (avl *AVL) Predecessor(v AVLInterface) *AVLNode {
	return nil
}

func (avl *AVL) RevSearch(pred func(v AVLInterface) bool) *AVLNode {
	var dfs func(node *AVLNode) *AVLNode
	dfs = func(node *AVLNode) *AVLNode {
		if node == nil {
			return nil
		}

		if pred(node.Data) {
			lo := dfs(node.Right)
			if lo == nil {
				return node
			}

			return lo
		} else {
			return dfs(node.Left)
		}
	}

	return dfs(avl.head)
}

func (avl *AVL) SearchDFS(pred func(v AVLInterface) bool) *AVLNode {
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

func (avl *AVL) Search(pred func(v AVLInterface) bool) *AVLNode {
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
