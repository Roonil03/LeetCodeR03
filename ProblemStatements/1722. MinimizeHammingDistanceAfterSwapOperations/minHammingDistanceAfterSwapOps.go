func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	dsu := NewDSU(n)
	for _, v := range allowedSwaps {
		dsu.Union(v[0], v[1])
	}
	comp := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		r := dsu.Find(i)
		if comp[r] == nil {
			comp[r] = make(map[int]int)
		}
		comp[r][source[i]]++
		comp[r][target[i]]--
	}
	res := 0
	for _, freq := range comp {
		for _, v := range freq {
			if v > 0 {
				res += v
			}
		}
	}
	return res
}

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p, r := make([]int, n), make([]int, n)
	for i := range p {
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

func (d *DSU) Union(x, y int) {
	x, y = d.Find(x), d.Find(y)
	if x == y {
		return
	}
	if d.rank[x] < d.rank[y] {
		d.parent[x] = y
	} else {
		d.parent[y] = x
		if d.rank[x] == d.rank[y] {
			d.rank[x]++
		}
	}
}