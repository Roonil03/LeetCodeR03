func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	et := NewEulerTour(n)
	for _, i := range edges {
		et.AddEdge(i[0], i[1])
	}
	et.Build(0)
	res := make([]bool, 0)
	ft := NewFenwick(n)
	ch := []byte(s)
	unc := func(u, v int) bool {
		return et.tin[u] <= et.tin[v] && et.tout[u] >= et.tout[v]
	}
	lca := func(u, v int) int {
		if unc(u, v) {
			return u
		}
		if unc(v, u) {
			return v
		}
		for i := sz - 1; i >= 0; i-- {
			if !unc(et.up[u][i], v) {
				u = et.up[u][i]
			}
		}
		return et.up[u][0]
	}
	h1 := func(u int, c byte) {
		m := 1 << (c - 'a')
		ft.Update(et.tin[u], m)
		ft.Update(et.tout[u]+1, m)
	}
	for i := 0; i < n; i++ {
		h1(i, ch[i])
	}
	for _, q := range queries {
		pars := strings.Fields(q)
		if pars[0] == "update" {
			u, _ := strconv.Atoi(pars[1])
			c := pars[2][0]
			dif := (1 << (ch[u] - 'a')) ^ (1 << (c - 'a'))
			ft.Update(et.tin[u], dif)
			ft.Update(et.tout[u]+1, dif)
			ch[u] = c
		} else {
			u, _ := strconv.Atoi(pars[1])
			v, _ := strconv.Atoi(pars[2])
			l := lca(u, v)
			m := ft.Query(et.tin[u]) ^ ft.Query(et.tin[v]) ^ (1 << (ch[l] - 'a'))
			res = append(res, bits.OnesCount(uint(m)) <= 1)
		}
	}
	return res
}

const sz = 17

type Fenwick struct {
	n    int
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{n: n, tree: make([]int, n+1)}
}

func (f *Fenwick) Update(i, v int) {
	for ; i <= f.n; i += i & -i {
		f.tree[i] ^= v
	}
}

func (f *Fenwick) Query(i int) int {
	sum := 0
	for ; i > 0; i -= i & -i {
		sum ^= f.tree[i]
	}
	return sum
}

type EulerTour struct {
	n         int
	adj       [][]int
	tin, tout []int
	timer     int
	up        [][]int
}

func NewEulerTour(n int) *EulerTour {
	up := make([][]int, n)
	for i := range up {
		up[i] = make([]int, 19)
	}
	return &EulerTour{n: n, adj: make([][]int, n), tin: make([]int, n), tout: make([]int, n), up: up}
}

func (e *EulerTour) AddEdge(u, v int) {
	e.adj[u] = append(e.adj[u], v)
	e.adj[v] = append(e.adj[v], u)
}

func (e *EulerTour) dfs(u, p int) {
	e.timer++
	e.tin[u] = e.timer
	e.up[u][0] = p
	if p == -1 {
		e.up[u][0] = u
	}
	for i := 1; i < 19; i++ {
		e.up[u][i] = e.up[e.up[u][i-1]][i-1]
	}
	for _, v := range e.adj[u] {
		if v != p {
			e.dfs(v, u)
		}
	}
	e.tout[u] = e.timer
}

func (e *EulerTour) Build(root int) {
	e.timer = 0
	e.dfs(root, -1)
}