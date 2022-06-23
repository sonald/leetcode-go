package main

import "container/list"

//BFS
func findBottomLeftValue(root *TreeNode) int {
	var stack = list.New()
	stack.PushBack(root)
	var np *TreeNode
	for stack.Len() > 0 {
		e := stack.Front()
		np = stack.Remove(e).(*TreeNode)

		if np.Right != nil {
			stack.PushBack(np.Right)
		}

		if np.Left != nil {
			stack.PushBack(np.Left)
		}
	}

	return np.Val
}

//BFS: this linear stack is faster than previous List stack
func findBottomLeftValue2(root *TreeNode) int {
	var stack = []*TreeNode{root}
	var np *TreeNode
	for len(stack) > 0 {
		np = stack[0]
		stack = stack[1:]
		if np.Right != nil {
			stack = append(stack, np.Right)
		}

		if np.Left != nil {
			stack = append(stack, np.Left)
		}
	}

	return np.Val
}

//DFS
func findBottomLeftValue3(root *TreeNode) int {
	maxd, res := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(np *TreeNode, d int) {
		if np == nil {
			return
		}
		if maxd < d && np.Left == nil && np.Right == nil {
			res = np.Val
			maxd = d
			return
		}

		dfs(np.Left, d+1)
		dfs(np.Right, d+1)
	}

	dfs(root, 1)
	return res
}
