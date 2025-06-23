func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	e1 := make([][]int, len(edges))
	for i, e := range edges {
		e1[i] = []int{e[0], e[1], e[2], i}
	}
	sort.Slice(e1, func(i, j int) bool {
		return e1[i][2] < e1[j][2]
	})
	kk := make([][]int, len(e1))
	for i, e := range e1 {
		kk[i] = []int{e[0], e[1], e[2]}
	}
	m := kruskal(n, kk, -1, -1)
	r1, r2 := []int{}, []int{}
	for i, e := range e1 {
		if kruskal(n, kk, i, -1) > m {
			r1 = append(r1, e[3])
		} else if kruskal(n, kk, -1, i) == m {
			r2 = append(r2, e[3])
		}
	}
	return [][]int{r1, r2}
}

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	xRoot, yRoot := d.Find(x), d.Find(y)
	if xRoot == yRoot {
		return false
	}
	if d.rank[xRoot] < d.rank[yRoot] {
		d.parent[xRoot] = yRoot
	} else if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[yRoot] = xRoot
	} else {
		d.parent[yRoot] = xRoot
		d.rank[xRoot]++
	}
	return true
}
func kruskal(n int, edges [][]int, bannedEdge int, forcedEdge int) int {
	dsu := NewDSU(n)
	cost := 0
	used := 0

	if forcedEdge != -1 {
		e := edges[forcedEdge]
		if dsu.Union(e[0], e[1]) {
			cost += e[2]
			used++
		}
	}
	for i, e := range edges {
		if i == bannedEdge {
			continue
		}
		if forcedEdge != -1 && i == forcedEdge {
			continue
		}
		if dsu.Union(e[0], e[1]) {
			cost += e[2]
			used++
			if used == n-1 {
				break
			}
		}
	}
	if used == n-1 {
		return cost
	}
	return 1 << 30
}