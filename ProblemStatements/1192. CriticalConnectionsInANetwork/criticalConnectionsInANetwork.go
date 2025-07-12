	func criticalConnections(n int, connections [][]int) [][]int {
		adj := make([][]int, n)
		for _, conn := range connections {
			u, v := conn[0], conn[1]
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		res := [][]int{}
		disc := make([]int, n)
		low := make([]int, n)
		time := 1
		var dfs func(u, par int)
		dfs = func(u, par int) {
			disc[u] = time
			low[u] = time
			time++
			for _, v := range adj[u] {
				if v == par {
					continue
				}
				if disc[v] == 0 {
					dfs(v, u)
					low[u] = min(low[u], low[v])
					if low[v] > disc[u] {
						res = append(res, []int{u, v})
					}
				} else {
					low[u] = min(low[u], disc[v])
				}
			}
		}
		dfs(0, -1)
		return res
	}