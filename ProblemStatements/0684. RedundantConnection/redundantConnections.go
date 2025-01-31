type UnionFind struct {
	parent, size []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{parent: make([]int, n), size: make([]int, n)}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) find(x int) int {
	if x == uf.parent[x] {
		return x
	}
	uf.parent[x] = uf.find(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) union(u, v int) bool {
	pu, pv := uf.find(u), uf.find(v)
	if pu == pv {
		return false
	}
	if uf.size[pu] > uf.size[pv] {
		uf.size[pu] += uf.size[pv]
		uf.parent[pv] = pu
	} else {
		uf.size[pv] += uf.size[pu]
		uf.parent[pu] = pv
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n)
	for _, e := range edges {
		if !uf.union(e[0]-1, e[1]-1) {
			return []int{e[0], e[1]}
		}
	}
	return nil
}
