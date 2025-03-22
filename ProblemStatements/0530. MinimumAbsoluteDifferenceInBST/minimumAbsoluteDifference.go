/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	_, _, res := helper(root)
	return res
}

func helper(n *TreeNode) (int, int, int) {
	m1, m2, dif := n.Val, n.Val, math.MaxInt32
	if n.Left != nil {
		l1, l2, ld := helper(n.Left)
		m1 = l1
		if t := n.Val - l2; t < dif {
			dif = t
		}
		if ld < dif {
			dif = ld
		}
	}
	if n.Right != nil {
		r1, r2, rd := helper(n.Right)
		m2 = r2
		if t := r1 - n.Val; t < dif {
			dif = t
		}
		if rd < dif {
			dif = rd
		}
	}
	return m1, m2, dif
}