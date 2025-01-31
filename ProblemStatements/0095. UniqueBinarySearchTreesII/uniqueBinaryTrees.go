/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return subs(1, n)
}

func subs(s, e int) []*TreeNode {
	res := []*TreeNode{}
	if s > e {
		res = append(res, nil)
		return res
	}
	for i := s; i <= e; i++ {
		lSubs := subs(s, i-1)
		rSubs := subs(i+1, e)
		for _, k := range lSubs {
			for _, j := range rSubs {
				temp := &TreeNode{Val: i}
				temp.Left = k
				temp.Right = j
				res = append(res, temp)
			}
		}
	}
	return res
}