package tree

import "math"

/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

链接：https://leetcode-cn.com/problems/validate-binary-search-tree
*/

func IsValidBSTWithInorderTravel(root *TreeNode) bool {
	if root == nil {
		return true
	}
	inorder := Inorder(root)
	if validSliceIncrease(inorder) {
		return true
	}
	return false
}
func Inorder(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var left, right []int
	if root.Left != nil {
		left = Inorder(root.Left)
	}
	if root.Right != nil {
		right = Inorder(root.Right)
	}
	res := append(left, root.Val)
	res = append(res, right...)
	return res
}

func validSliceIncrease(nums []int) bool {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] { //不可以相等，package sort 里面的判断如果相等也会返回true
			return false
		}
	}
	return true
}

func IsValidBSTWithRecrusion(root *TreeNode) bool {
	return validWithRecursion(root, math.MinInt64, math.MaxInt64)
}
func validWithRecursion(root *TreeNode, min, max int) bool {
	//递归结束条件
	if root == nil {
		return true
	}
	// 判断节点的值是不是在区间呢，不是的话就false结束
	if root.Val <= min || root.Val >= max {
		return false
	}
	return validWithRecursion(root.Left, min, root.Val) && validWithRecursion(root.Right, root.Val, max)
}
