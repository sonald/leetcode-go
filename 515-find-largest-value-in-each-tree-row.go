package main

import "math"

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var sp = &TreeNode{} // level separator
	var stack = []*TreeNode{root, sp}

	var maxv = math.MinInt32
	for len(stack) > 0 {
		np := stack[0]
		stack = stack[1:]
		if np == sp {
			res = append(res, maxv)
			if len(stack) == 0 {
				break
			}
			stack = append(stack, sp)
			maxv = math.MinInt32
			continue
		}

		if maxv < np.Val {
			maxv = np.Val
		}

		if np.Left != nil {
			stack = append(stack, np.Left)
		}

		if np.Right != nil {
			stack = append(stack, np.Right)
		}

	}

	return res
}

func largestValues2(root *TreeNode) []int {
	var d = map[int]int{}
	var dfs func(*TreeNode, int)
	dfs = func(np *TreeNode, level int) {
		if np == nil {
			return
		}
		if v, ok := d[level]; !ok {
			d[level] = np.Val
		} else {
			if v < np.Val {
				d[level] = np.Val
			}
		}

		dfs(np.Left, level+1)
		dfs(np.Right, level+1)
	}

	dfs(root, 0)

	var res = make([]int, len(d))
	for l, v := range d {
		res[l] = v
	}
	return res
}
