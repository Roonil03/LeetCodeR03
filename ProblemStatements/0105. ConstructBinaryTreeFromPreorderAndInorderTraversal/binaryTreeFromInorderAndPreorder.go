/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	pi := 0
	var bt func(in []int, pre []int, st int, en int) *TreeNode
	bt = func(in []int, pre []int, st int, en int) *TreeNode {
		if st > en {
			return nil
		}
		r := &TreeNode{Val: pre[pi]}
		pi++
		if st == en {
			return r
		}
		idx := st
		for i := st; i <= en; i++ {
			if in[i] == r.Val {
				idx = i
				break
			}
		}
		r.Left = bt(in, pre, st, idx-1)
		r.Right = bt(in, pre, idx+1, en)
		return r
	}
	return bt(inorder, preorder, 0, len(inorder)-1)
}