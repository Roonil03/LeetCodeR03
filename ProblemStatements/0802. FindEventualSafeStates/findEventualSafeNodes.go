func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	rev := make([][]int, n)
	deg := make([]int, n)
	for i, nei := range graph {
		for _, n1 := range nei {
			rev[n1] = append(rev[n1], i)
		}
		deg[i] = len(nei)
	}
	q := []int{}
	res := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = append(res, node)
		for _, p := range rev[node] {
			deg[p]--
			if deg[p] == 0 {
				q = append(q, p)
			}
		}
	}
	sort.Ints(res)
	return res
}