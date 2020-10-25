package tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	nums := []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	p, q, r := &TreeNode{Val: 2}, &TreeNode{Val: 8}, &TreeNode{Val: 4}

	tree := BuildTree(nums)
	ancestor := lowestCommonAncestor(tree, p, q)
	if ancestor.Val != 6 {
		t.Error("the ancestor returns ", ancestor.Val)
	}
	ancestor = lowestCommonAncestor(tree, p, r)
	if ancestor.Val != 2 {
		t.Error("the ancestor returns ", ancestor.Val)
	}

}
