package tree

import "testing"
import "algorithm/util"

func TestDFS(t *testing.T) {
	var nums []interface{} = []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	var root *TreeNode = BuildTree(nums)
	result := DFS(root)
	if !util.CheckSliceEq(result, []int{6, 2, 0, 4, 3, 5, 8, 7, 9}) {
		t.Error("the DFS() returns wrong answer:", result)
	}
}
