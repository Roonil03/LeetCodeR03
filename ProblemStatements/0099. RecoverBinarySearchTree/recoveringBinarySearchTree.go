/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(rt *TreeNode) {
	var f, s, pr, pt *TreeNode
	cr := rt
	for cr != nil {
		if cr.Left == nil {
			if pr != nil && pr.Val > cr.Val {
				if f == nil {
					f = pr
				}
				s = cr
			}
			pr = cr
			cr = cr.Right
			continue
		}
		pt = cr.Left
		for pt.Right != nil && pt.Right != cr {
			pt = pt.Right
		}
		if pt.Right == nil {
			pt.Right = cr
			cr = cr.Left
		} else {
			pt.Right = nil
			if pr != nil && pr.Val > cr.Val {
				if f == nil {
					f = pr
				}
				s = cr
			}
			pr = cr
			cr = cr.Right
		}
	}
	if f != nil && s != nil {
		f.Val, s.Val = s.Val, f.Val
	}
}