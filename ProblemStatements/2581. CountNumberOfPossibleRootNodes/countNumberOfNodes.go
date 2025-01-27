var res int

func rootCount(edg [][]int, gus [][]int, k int) int {
	res = 0
	adj := make([][]int, len(edg)+1)
	for _, e := range edg {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	vis := map[int]bool{0: true}
	gSet := make(map[pair]bool)
	for _, g := range gus {
		gSet[pair{f: g[0], s: g[1]}] = true
	}
	cnt := findCnt(0, adj, vis, gSet)
	vis = map[int]bool{0: true}
	traverse(0, cnt, 0, k, adj, vis, gSet)
	return res
}

func findCnt(n int, adj [][]int, vis map[int]bool, gSet map[pair]bool) int {
	var cnt int
	for _, nei := range adj[n] {
		if !vis[nei] {
			vis[nei] = true
			if gSet[pair{f: n, s: nei}] {
				cnt++
			}
			cnt += findCnt(nei, adj, vis, gSet)
		}
	}
	return cnt
}

func traverse(n, cnt, rev, k int, adj [][]int, vis map[int]bool, gSet map[pair]bool) {
	for _, nei := range adj[n] {
		if !vis[nei] {
			vis[nei] = true
			r, c := rev, cnt
			if gSet[pair{f: nei, s: n}] {
				r++
			}
			if gSet[pair{f: n, s: nei}] {
				c--
			}
			traverse(nei, c, r, k, adj, vis, gSet)
		}
	}
	if cnt+rev >= k {
		res++
	}
}

type pair struct {
	f int
	s int
}