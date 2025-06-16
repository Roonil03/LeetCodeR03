/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, m int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val >= m {
			res = 1
			if node.Val > m {
				m = node.Val
			}
		}
		return res + dfs(node.Left, m) + dfs(node.Right, m)
	}
	return dfs(root, root.Val)
}