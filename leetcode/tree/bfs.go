package tree

import "container/list"

func BFS(root *TreeNode) []int {
	var result []int
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		node := e.Value.(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
	return result
}
