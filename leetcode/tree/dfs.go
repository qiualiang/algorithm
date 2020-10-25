package tree

//递归
var visited []int

func DFS(root *TreeNode) []int {
	if root != nil {
		visited = append(visited, root.Val)
	}
	if root.Left != nil {
		DFS(root.Left)
	}
	if root.Right != nil {
		DFS(root.Right)
	}
	return visited
}
