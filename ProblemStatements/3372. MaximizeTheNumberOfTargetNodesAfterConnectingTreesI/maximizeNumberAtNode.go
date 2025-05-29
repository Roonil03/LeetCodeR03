func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1
	g1 := buildGraph(edges1, n)
	g2 := buildGraph(edges2, m)
	if k == 0 {
		ans := make([]int, n)
		for i := range ans {
			ans[i] = 1
		}
		return ans
	}
	cnt1 := bfsCount(g1, k)
	cnt2 := bfsCount(g2, k-1)
	best2 := 0
	for _, x := range cnt2 {
		if x > best2 {
			best2 = x
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = cnt1[i] + best2
	}
	return ans
}

func buildGraph(edges [][]int, sz int) [][]int {
	g := make([][]int, sz)
	for i := range g {
		g[i] = []int{}
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	return g
}

func bfsCount(g [][]int, K int) []int {
	N := len(g)
	result := make([]int, N)
	visited := make([]int, N)
	stamp := 1
	queue := make([]int, 0, N)
	for start := 0; start < N; start, stamp = start+1, stamp+1 {
		if K < 0 {
			result[start] = 1
			continue
		}
		queue = queue[:0]
		queue = append(queue, start)
		visited[start] = stamp
		head, tail, level, count := 0, 1, 0, 1
		for head < tail && level < K {
			sz := tail - head
			for i := 0; i < sz; i++ {
				u := queue[head]
				head++
				for _, v := range g[u] {
					if visited[v] != stamp {
						visited[v] = stamp
						queue = append(queue, v)
						tail++
						count++
					}
				}
			}
			level++
		}
		result[start] = count
	}
	return result
}