/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	n := make(map[int]*TreeNode)
	child := make(map[int]bool)
	for _, v := range descriptions {
		p, c, fg := v[0], v[1], v[2]
		if n[p] == nil {
			n[p] = &TreeNode{Val: p}
		}
		if n[c] == nil {
			n[c] = &TreeNode{Val: c}
		}
		if fg == 1 {
			n[p].Left = n[c]
		} else {
			n[p].Right = n[c]
		}
		child[c] = true
	}
	for _, v := range descriptions {
		if !child[v[0]] {
			return n[v[0]]
		}
	}
	return nil
}