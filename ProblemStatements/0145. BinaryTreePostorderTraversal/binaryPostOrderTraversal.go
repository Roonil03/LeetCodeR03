/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var trav func(*TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		trav(root.Left)
		trav(root.Right)
		res = append(res, root.Val)
		return
	}
	trav(root)
	return res
}