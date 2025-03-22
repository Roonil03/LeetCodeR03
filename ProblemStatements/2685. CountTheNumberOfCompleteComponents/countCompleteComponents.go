func countCompleteComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	vis := make([]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			ver := []int{}
			dfs(i, adj, vis, &ver)
			if comp(ver, adj) {
				res++
			}
		}
	}
	return res
}

func dfs(node int, adj [][]int, vis []bool, ver *[]int) {
	vis[node] = true
	*ver = append(*ver, node)
	for _, neigh := range adj[node] {
		if !vis[neigh] {
			dfs(neigh, adj, vis, ver)
		}
	}
}

func comp(nodes []int, adj [][]int) bool {
	// n := len(vertices)
	// expectedEdges := (n * (n - 1)) / 2
	// vertexSet := make(map[int]bool)
	// for _, v := range vertices {
	//     vertexSet[v] = true
	// }
	// edgeCount := 0
	// counted := make(map[string]bool)
	// for _, v := range vertices {
	//     for _, neighbor := range adj[v] {
	//         if !vertexSet[neighbor] {
	//             return false
	//         }
	//         edge := ""
	//         if v < neighbor {
	//             edge = fmt.Sprintf("%d-%d", v, neighbor)
	//         } else {
	//             edge = fmt.Sprintf("%d-%d", neighbor, v)
	//         }
	//         if !counted[edge] {
	//             edgeCount++
	//             counted[edge] = true
	//         }
	//     }
	// }
	// return edgeCount == expectedEdges
	size := len(nodes)
	for _, node := range nodes {
		if len(adj[node]) != size-1 {
			return false
		}
	}
	return true
}