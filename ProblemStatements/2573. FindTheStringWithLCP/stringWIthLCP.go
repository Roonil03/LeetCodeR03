func findTheString(lcp [][]int) string {
	n := len(lcp)
	dsu := NewDSU(n)
	for i := range n {
		if len(lcp[i]) != n {
			return ""
		}
		for j := i + 1; j < n; j++ {
			if lcp[i][j] > 0 {
				dsu.Union(i, j)
			}
		}
	}
	w := make([]byte, n)
	a := make(map[int]byte, n)
	b := byte('a')
	for i := range n {
		root := dsu.Find(i)
		ch, ex := a[root]
		if !ex {
			if b > 'z' {
				return ""
			}
			ch = b
			a[root] = ch
			b++
		}
		w[i] = ch
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		p := 0
		for j := n - 1; j >= 0; j-- {
			old := dp[j]
			cur := 0
			if w[i] == w[j] {
				cur = 1 + p
			}
			if cur != lcp[i][j] {
				return ""
			}
			dp[j] = cur
			p = old
		}
		dp[n] = 0
	}
	return string(w)
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