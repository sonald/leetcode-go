package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

    i := 0
	res := nums[0]
    for k, v := range nums[0:] {
		if res < v {
			res = v
            i = k
		}
	}

    return &TreeNode{
        Val: res,
        Left: constructMaximumBinaryTree(nums[:i]),
        Right: constructMaximumBinaryTree(nums[i+1:]),
    }
}


//this is actually the implicit monotone-stack algorithm
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

    head := &TreeNode{Val: nums[0]}

    for _, v := range nums[1:] {
        np := &TreeNode{Val: v}
        if v < head.Val {
            r := &head.Right
            for {
                if *r == nil {
                    *r = np
                    break
                }

                if np.Val > (*r).Val {
                    *r, np.Left = np, *r
                    break
                }
                r = &(*r).Right
            }

        } else {
            head, np.Left = np, head
        }
    }

    return head
}
