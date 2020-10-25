package tree

import "testing"
import "algorithm/util"

func TestBFS(t *testing.T) {
	var nums []interface{} = []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	var root *TreeNode = BuildTree(nums)
	result := BFS(root)
	if !util.CheckSliceEq(result, []int{6, 2, 8, 0, 4, 7, 9, 3, 5}) {
		t.Error("the BFS() returns wrong answer:", result)
	}
}
