/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(n *Node) *Node {
	if n == nil {
		return nil
	}
	q := []*Node{n}
	m := map[int]*Node{n.Val: &Node{Val: n.Val}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		c := m[v.Val]
		for _, u := range v.Neighbors {
			if _, ok := m[u.Val]; !ok {
				m[u.Val] = &Node{Val: u.Val}
				q = append(q, u)
			}
			c.Neighbors = append(c.Neighbors, m[u.Val])
		}
	}
	return m[n.Val]
}