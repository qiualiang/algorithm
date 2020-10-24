package btree

import "testing"

func TestInorder1(t *testing.T) {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	if !IsValidBSTWithInorderTravel(root) {
		t.Fatal("It is a binary tree!")
	}
}
func TestInorder2(t *testing.T) {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	if !IsValidBSTWithRecrusion(root) {
		t.Fatal("It is a binary tree!")
	}
}
