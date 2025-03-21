/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := []int{}
	var dfs func(n *TreeNode, lvl int)
	dfs = func(n *TreeNode, lvl int) {
		if n == nil {
			return
		}
		if len(res) == lvl {
			res = append(res, n.Val)
		}
		dfs(n.Right, lvl+1)
		dfs(n.Left, lvl+1)
		return
	}
	dfs(root, 0)
	return res
}