/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var get func(*TreeNode) []int
	var dfs func(*TreeNode, *[]int)
	dfs = func(node *TreeNode, l *[]int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			*l = append(*l, node.Val)
			return
		}
		dfs(node.Left, l)
		dfs(node.Right, l)
	}
	get = func(root *TreeNode) []int {
		l := []int{}
		dfs(root, &l)
		return l
	}
	l1, l2 := get(root1), get(root2)
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}