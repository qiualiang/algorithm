package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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

func main() {

	nums := []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	p, q, r := &TreeNode{Val: 2}, &TreeNode{Val: 4}, &TreeNode{Val: 8}
	var root *TreeNode
	tree := createTree(root, nums, len(nums), 0)
	ancestor := lowestCommonAncestor(tree, p, q)
	fmt.Println("ancestor:", ancestor.Val)
	ancestor = lowestCommonAncestor(tree, p, r)
	fmt.Println("ancestor:", ancestor.Val)

	// print(tree)

}
func createTree(root *TreeNode, nums []interface{}, len int, index int) (tree *TreeNode) {
	if index >= len {
		return
	}
	if val, ok := nums[index].(int); ok {
		root = new(TreeNode)
		root.Val = val
	} else {
		root = nil
	}
	if root != nil {
		root.Left = createTree(root.Left, nums, len, 2*index+1)
		root.Right = createTree(root.Right, nums, len, 2*index+2)
	}
	return root
}

func print(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Val)
	fmt.Print(" ")
	print(tree.Left)
	print(tree.Right)
}
