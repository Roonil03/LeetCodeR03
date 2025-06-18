/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	a, b, res := root.Val, 1, 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		ls := 0
		for i := 0; i < n; i++ {
			n := q[0]
			q = q[1:]
			ls += n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if ls > a {
			a = ls
			res = b
		}
		b++
	}
	return res
}