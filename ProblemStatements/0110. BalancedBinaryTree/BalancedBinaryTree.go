/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, res := h(root)
	return res
}

func h(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}
	l, bl := h(n.Left)
	r, br := h(n.Right)
	if math.Abs(float64(l-r)) > 1 {
		return -1, false
	}
	return int(math.Max(float64(l), float64(r)) + 1), bl && br
}