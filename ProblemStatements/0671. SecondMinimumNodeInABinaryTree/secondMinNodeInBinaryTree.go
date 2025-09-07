/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	a := root.Val
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return -1
		}
		if n.Val > a {
			return n.Val
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		if l == -1 {
			return r
		}
		if r == -1 {
			return l
		}
		return min(l, r)
	}
	return dfs(root)
}