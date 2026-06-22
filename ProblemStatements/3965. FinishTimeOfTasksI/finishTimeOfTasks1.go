func finishTime(n int, edges [][]int, baseTime []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	var dfs func(int) int64
	dfs = func(u int) int64 {
		if len(adj[u]) == 0 {
			return int64(baseTime[u])
		}
		a, b := int64(-1<<63), int64(1<<63-1)
		for _, v := range adj[u] {
			f := dfs(v)
			if f > a {
				a = f
			}
			if f < b {
				b = f
			}
		}
		return 2*a - b + int64(baseTime[u])
	}
	return dfs(0)
}