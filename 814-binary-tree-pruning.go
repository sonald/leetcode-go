package main

func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(np *TreeNode) *TreeNode {
		if np == nil {
			return np
		}

		np.Left = dfs(np.Left)
		np.Right = dfs(np.Right)

		if np.Left == nil && np.Right == nil && np.Val == 0 {
			return nil
		}

		return np
	}

	return dfs(root)
}
