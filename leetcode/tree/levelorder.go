package tree

/* 102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

*/
var results [][]int

func levelOrder(root *TreeNode) [][]int {
	results = nil
	if root == nil {
		return nil
	}
	travel([]*TreeNode{root})
	return results
}
func travel(level []*TreeNode) {
	if len(level) == 0 {
		return
	}
	var nextLevel []*TreeNode
	var values []int = make([]int, len(level))
	for i, node := range level {
		values[i] = node.Val
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}
	results = append(results, values)
	travel(nextLevel)
}

//Using dfs to get
// func levelOrderDFS(root *TreeNode) [][]int {

// }
