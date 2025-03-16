/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func maxPathSum(root *TreeNode) int {
//     res := math.MinInt32
//     helper(root, &res)
//     return res
// }

// func helper(node *TreeNode, s *int)int{
//     if node == nil{
//         return 0
//     }
//     l := max(helper(node.Left, s), 0)
//     r := max(helper(node.Right, s), 0)
//     res := node.Val + l + r
//     *s = max(*s, res)
//     return node.Val + max(l + r)
// }

func maxPathSum(root *TreeNode) int {
	globalMax := -1 << 63
	dfs(root, &globalMax)
	return globalMax
}

func dfs(root *TreeNode, globalMax *int) int {
	if root == nil {
		return 0
	}
	pathSumFromLeft := max(dfs(root.Left, globalMax), 0)
	pathSumFromRight := max(dfs(root.Right, globalMax), 0)
	*globalMax = max(*globalMax, root.Val+pathSumFromLeft+pathSumFromRight)
	return root.Val + max(pathSumFromLeft, pathSumFromRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}