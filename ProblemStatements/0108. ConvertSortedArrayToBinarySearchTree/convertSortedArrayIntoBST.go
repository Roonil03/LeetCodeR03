/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return help(nums, 0, len(nums)-1)
}

func help(nums []int, a, b int) *TreeNode {
	if a > b {
		return nil
	}
	mid := a + (b-a)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = help(nums, a, mid-1)
	root.Right = help(nums, mid+1, b)
	return root
}