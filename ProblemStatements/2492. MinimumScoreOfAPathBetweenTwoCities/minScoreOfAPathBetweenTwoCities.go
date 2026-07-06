func minScore(n int, roads [][]int) int {
	g := make([][][2]int, n+1)
	for _, r := range roads {
		g[r[0]] = append(g[r[0]], [2]int{r[1], r[2]})
		g[r[1]] = append(g[r[1]], [2]int{r[0], r[2]})
	}
	vis := make([]bool, n+1)
	vis[1] = true
	q := []int{1}
	res := int(1<<31 - 1)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range g[cur] {
			res = min(res, v[1])
			// fmt.Println(res)
			if !vis[v[0]] {
				vis[v[0]] = true
				q = append(q, v[0])
			}
		}
	}
	return res
}