func dfs(node int, adj [][]int, visited []bool, group *[]int) {
	visited[node] = true
	*group = append(*group, node)
	for _, nei := range adj[node] {
		if !visited[nei] {
			dfs(nei, adj, visited, group)
		}
	}
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	adj := make([][]int, 26)
	for i := 0; i < len(s1); i++ {
		u := int(s1[i] - 'a')
		v := int(s2[i] - 'a')
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	visited := make([]bool, 26)
	rep := make([]byte, 26)
	for i := 0; i < 26; i++ {
		if !visited[i] {
			group := []int{}
			dfs(i, adj, visited, &group)
			minChar := byte('z')
			for _, idx := range group {
				if ch := byte(idx + 'a'); ch < minChar {
					minChar = ch
				}
			}
			for _, idx := range group {
				rep[idx] = minChar
			}
		}
	}
	res := []byte{}
	for i := 0; i < len(baseStr); i++ {
		ch := baseStr[i] - 'a'
		if rep[ch] == 0 {
			res = append(res, baseStr[i])
		} else {
			res = append(res, rep[ch])
		}
	}
	return string(res)
}