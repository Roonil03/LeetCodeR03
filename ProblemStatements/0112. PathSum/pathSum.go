/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if help(root, 0, targetSum) {
		return true
	}
	return false
}

func help(root *TreeNode, cur, aim int) bool {
	sum := cur + root.Val
	if sum == aim && root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return false
	}
	if root.Left != nil {
		if help(root.Left, sum, aim) {
			return true
		}
	}
	if root.Right != nil {
		if help(root.Right, sum, aim) {
			return true
		}
	}
	return false
}