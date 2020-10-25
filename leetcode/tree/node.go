package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
root = [6,2,8,0,4,7,9,nil,nil,3,5]
一层一层生成二叉树
*/
func BuildTree(nums []interface{}) *TreeNode {
	var root *TreeNode
	tree := createTree(root, nums, len(nums), 0)
	return tree
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

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var left, right []int
	if root.Left != nil {
		left = preorder(root.Left)
	}
	if root.Right != nil {
		right = preorder(root.Right)
	}
	res := append([]int{root.Val}, left...)
	res = append(res, right...)
	return res
}
func inorder(root *TreeNode) []int {
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
func postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var left, right []int
	if root.Left != nil {
		left = postorder(root.Left)
	}
	if root.Right != nil {
		right = postorder(root.Right)
	}
	res := append(left, right...)
	res = append(res, root.Val)
	return res
}
