/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// if root.Left == nil && root.Right == nil{
	//     return 1
	// }
	// if root.Left == nil{
	//     return minDepth(root.Right)+1
	// }
	// if root.Right == nil{
	//     return minDepth(root.Left)+1
	// }
	// return min(minDepth(root.Left)+1, minDepth(root.Right)+1)
	q := []*TreeNode{root}
	res := 1
	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			n := q[0]
			q = q[1:]
			if n.Left == nil && n.Right == nil {
				return res
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res++
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}