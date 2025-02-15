/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	var dfs func(n *TreeNode, cs int)

	dfs = func(n *TreeNode, cs int) {
		if n == nil {
			return
		}
		path = append(path, n.Val)
		cs += n.Val
		if n.Left == nil && n.Right == nil && cs == targetSum {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		dfs(n.Left, cs)
		dfs(n.Right, cs)
		path = path[:len(path)-1]
	}
	dfs(root, 0)
	return res
}