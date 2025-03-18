/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := helperLeft(root)
	rh := helperRight(root)
	if lh == rh {
		return (1 << lh) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func helperLeft(n *TreeNode) int {
	h := 0
	for n != nil {
		h++
		n = n.Left
	}
	return h
}

func helperRight(n *TreeNode) int {
	h := 0
	for n != nil {
		h++
		n = n.Right
	}
	return h
}