func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0
	q := make([]int, 0, n)
	q = append(q, 1)
	head := 0
	dep := 0
	for head < len(q) {
		u := q[head]
		head++
		dep = dist[u]
		for _, v := range adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				q = append(q, v)
			}
		}
	}
	if dep == 0 {
		return 0
	}
	res := 1
	base, exp, mod := 2, dep-1, int(1e9+7)
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return res
}