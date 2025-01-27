/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	//var nums = []int{}
	if root == nil {
		return []int{}
	}
	nums := append(inorderTraversal(root.Left), root.Val)
	return append(nums, inorderTraversal(root.Right)...)
}