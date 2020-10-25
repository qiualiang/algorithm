package tree

import (
	"algorithm/util"
	"testing"
)

var nums []interface{} = []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
var pre []int = []int{6, 2, 0, 4, 3, 5, 8, 7, 9}
var in []int = []int{0, 2, 3, 4, 5, 6, 7, 8, 9}
var post []int = []int{0, 3, 5, 4, 2, 7, 9, 8, 6}
var root *TreeNode = BuildTree(nums)

func TestPreorder(t *testing.T) {
	if !util.CheckSliceEq(preorder(root), pre) {
		t.Error("pre order error :", preorder(root))
	}

}
func TestInorder(t *testing.T) {
	if !util.CheckSliceEq(inorder(root), in) {
		t.Error("in order error ", inorder(root))
	}
}
func TestPostorder(t *testing.T) {
	if !util.CheckSliceEq(postorder(root), post) {
		t.Error("post order error ", postorder(root))
	}
}
