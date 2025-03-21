/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
	memo := make(map[int][]*TreeNode)
	return helper(n, memo)
}

func helper(n int, memo map[int][]*TreeNode) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	if t, exists := memo[n]; exists {
		return t
	}
	res := []*TreeNode{}
	for l := 1; l < n; l += 2 {
		r := n - 1 - l
		ls := helper(l, memo)
		rs := helper(r, memo)
		for _, i := range ls {
			for _, j := range rs {
				root := &TreeNode{
					Val:   0,
					Left:  i,
					Right: j,
				}
				res = append(res, root)
			}
		}
	}
	memo[n] = res
	return res
}