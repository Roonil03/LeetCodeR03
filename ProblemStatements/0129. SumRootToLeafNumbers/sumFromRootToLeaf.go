/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(n *TreeNode, s int) int {
	if n == nil {
		return 0
	}
	s = s*10 + n.Val
	if n.Left == nil && n.Right == nil {
		return s
	}
	return dfs(n.Left, s) + dfs(n.Right, s)
}