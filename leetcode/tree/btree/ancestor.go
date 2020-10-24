package btree

/*
235.给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findPorQ(root, p, q)
}
func findPorQ(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	if right != nil && left != nil {
		return root
	}
	return nil
}
