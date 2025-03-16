/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
	} else if root.Left == nil {
		root.Left = root.Right
		root.Right = nil
	} else {
		root.Right = root.Left
		root.Left = nil
	}
	helper(root.Left)
	helper(root.Right)
	return
}