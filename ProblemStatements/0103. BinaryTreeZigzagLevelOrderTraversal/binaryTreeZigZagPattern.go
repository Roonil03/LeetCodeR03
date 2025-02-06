/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	lr := true
	for len(q) > 0 {
		size := len(q)
		row := make([]int, size)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			index := i
			if !lr {
				index = size - 1 - i
			}
			row[index] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		lr = !lr
		res = append(res, row)
	}
	return res
}