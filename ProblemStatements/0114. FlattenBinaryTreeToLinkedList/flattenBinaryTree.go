/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root
	for temp != nil {
		if temp.Left != nil {
			pre := temp.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = temp.Right
			temp.Right = temp.Left
			temp.Left = nil
		}
		temp = temp.Right
	}
	return
}