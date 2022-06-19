package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeSums(np *TreeNode, sums map[int]int) int {
	var l, r int
	if np.Left != nil {
		l = subtreeSums(np.Left, sums)
		sums[l]++
	}

	if np.Right != nil {
		r = subtreeSums(np.Right, sums)
		sums[r]++
	}

	return np.Val + l + r
}

func findFrequentTreeSum2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var sums = map[int]int{}
	v := subtreeSums(root, sums)
	sums[v]++

	var rev = map[int][]int{}
	var c = 0
	for k, f := range sums {
		rev[f] = append(rev[f], k)
		if c < f {
			c = f
		}
	}

	return rev[c]
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var sums = map[int]int{}
	var c = 0
	var subtreeSums func(np *TreeNode) int
	subtreeSums = func(np *TreeNode) int {
		if np == nil {
			return 0
		}

		v := np.Val + subtreeSums(np.Left) + subtreeSums(np.Right)
		sums[v]++
		if sums[v] > c {
			c = sums[v]
		}

		return v
	}
	subtreeSums(root)

	var res []int
	for s, f := range sums {
		if f == c {
			res = append(res, s)
		}
	}

	return res
}
