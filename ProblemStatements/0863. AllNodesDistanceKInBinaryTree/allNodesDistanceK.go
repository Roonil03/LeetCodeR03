/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	m := make(map[*TreeNode]*TreeNode)
	dfs(root, nil, m)
	q := []*TreeNode{target}
	vis := make(map[*TreeNode]bool)
	vis[target] = true
	d := 0
	for len(q) > 0 && d < k {
		qs := len(q)
		for i := 0; i < qs; i++ {
			cur := q[0]
			q = q[1:]
			can := []*TreeNode{
				cur.Left,
				cur.Right,
				m[cur],
			}
			for _, c := range can {
				if c != nil && !vis[c] {
					vis[c] = true
					q = append(q, c)
				}
			}
		}
		d++
	}
	res := make([]int, 0, len(q))
	for _, node := range q {
		res = append(res, node.Val)
	}
	return res
}

func dfs(node *TreeNode, p *TreeNode, m map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	m[node] = p
	dfs(node.Left, node, m)
	dfs(node.Right, node, m)
}