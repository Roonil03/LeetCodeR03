func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	up := make([][20]int, n+1)
	dep := make([]int, n+1)
	var dfs func(int, int, int)
	dfs = func(u, p, d int) {
		dep[u] = d
		up[u][0] = p
		for i := 1; i < 20; i++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}
		for _, v := range adj[u] {
			if v != p {
				dfs(v, u, d+1)
			}
		}
	}
	dfs(1, 0, 0)
	lca := func(u, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := 0; i < 20; i++ {
			if ((dep[u] - dep[v]) & (1 << i)) != 0 {
				u = up[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := 19; i >= 0; i-- {
			if up[u][i] != up[v][i] {
				u = up[u][i]
				v = up[v][i]
			}
		}
		return up[u][0]
	}
	p2 := make([]int, n+1)
	p2[0] = 1
	for i := 1; i <= n; i++ {
		p2[i] = (p2[i-1] * 2) % int(1e9+7)
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		if q[0] == q[1] {
			res[i] = 0
		} else {
			dist := dep[q[0]] + dep[q[1]] - 2*dep[lca(q[0], q[1])]
			res[i] = p2[dist-1]
		}
	}
	return res
}