/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(n *TreeNode) (int, int) {
	if n == nil {
		return math.MinInt, 0
	}
	lm, lc := dfs(n.Left)
	rm, rc := dfs(n.Right)
	m := max(n.Val, max(lm, rm))
	c := lc + rc
	if n.Val == m {
		c++
	}
	return m, c
}

func countDominantNodes(root *TreeNode) int {
	_, res := dfs(root)
	return res
}