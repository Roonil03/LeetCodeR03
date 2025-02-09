/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	a := postorder[len(postorder)-1]
	r := &TreeNode{Val: a}
	ina := 0
	for i, val := range inorder {
		if val == a {
			ina = i
			break
		}
	}
	r.Left = buildTree(inorder[:ina], postorder[:ina])
	r.Right = buildTree(inorder[ina+1:], postorder[ina:len(postorder)-1])
	return r
}