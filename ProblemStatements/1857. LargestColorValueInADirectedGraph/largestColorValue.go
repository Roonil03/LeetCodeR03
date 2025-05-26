func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	g := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		indeg[v]++
	}
	freq := make([][]int, n)
	for i := 0; i < n; i++ {
		freq[i] = make([]int, 26)
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	vis := 0
	res := 0
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		vis++
		id := int(colors[u] - 'a')
		freq[u][id]++
		if freq[u][id] > res {
			res = freq[u][id]
		}
		for _, v := range g[u] {
			for c := 0; c < 26; c++ {
				if freq[u][c] > freq[v][c] {
					freq[v][c] = freq[u][c]
				}
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if vis < n {
		return -1
	}
	return res
}