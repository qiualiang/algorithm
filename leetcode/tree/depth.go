package tree

/*104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
*/
func minDepth(root *TreeNode) int {
	//terminate
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	//divide and conquer
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)
	result := 1 + min(leftMinDepth, rightMinDepth)
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
