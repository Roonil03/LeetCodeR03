// Could not solve during contest time, has been solved post the contest

// func findMedian(n int, edges [][]int, queries [][]int) []int {
//     g := make([][][2]int, n)
//     for _, e := range edges {
//         u, v, w := e[0], e[1], e[2]
//         g[u] = append(g[u], [2]int{v, w})
//         g[v] = append(g[v], [2]int{u, w})
//     }
//     p := make([]int, n)
//     d := make([]int, n)
//     ds := make([]int, n)
//     st := []int{0}
//     vs := make([]bool, n)
//     vs[0] = true
//     for len(st) > 0 {
//         u := st[len(st)-1]
//         st = st[:len(st)-1]
//         for _, nw := range g[u] {
//             v, w := nw[0], nw[1]
//             if vs[v] {
//                 continue
//             }
//             vs[v] = true
//             d[v] = d[u] + 1
//             ds[v] = ds[u] + w
//             p[v] = u
//             st = append(st, v)
//         }
//     }
//     lg := bits.Len(uint(n))
//     pa := make([][]int, lg)
//     for i := range pa {
//         pa[i] = make([]int, n)
//         for j := range pa[i] {
//             pa[i][j] = -1
//         }
//     }
//     copy(pa[0], p)
//     for i := 1; i < lg; i++ {
//         for j := 0; j < n; j++ {
//             if pa[i-1][j] != -1 {
//                 pa[i][j] = pa[i-1][pa[i-1][j]]
//             }
//         }
//     }
//     lca := func(a, b int) int {
//         if d[a] < d[b] {
//             a, b = b, a
//         }
//         diff := d[a] - d[b]
//         for i := lg - 1; i >= 0; i-- {
//             if diff&(1<<i) != 0 {
//                 a = pa[i][a]
//             }
//         }
//         if a == b {
//             return a
//         }
//         for i := lg - 1; i >= 0; i-- {
//             if pa[i][a] != pa[i][b] {
//                 a = pa[i][a]
//                 b = pa[i][b]
//             }
//         }
//         return pa[0][a]
//     }
//     ans := make([]int, len(queries))
//     for i, q := range queries {
//         u, v := q[0], q[1]
//         l := lca(u, v)
//         tt := ds[u] + ds[v] - 2*ds[l]
//         hf := float64(tt) / 2.0
//         cr := u
//         cw := 0.0
//         for cr != l {
//             if cw >= hf {
//                 ans[i] = cr
//                 break
//             }
//             cr = pa[0][cr]
//             cw = float64(ds[u] - ds[cr])
//         }
//         if cr == l && cw < hf {
//             if cw >= hf {
//                 ans[i] = l
//             } else {
//                 cr = v
//                 for cr != l {
//                     tw := float64(ds[u]-ds[l]) + float64(ds[cr]-ds[l])
//                     if tw >= hf {
//                         ans[i] = cr
//                         break
//                     }
//                     cr = pa[0][cr]
//                 }
//                 if cr == l {
//                     ans[i] = l
//                 }
//             }
//         } else if cr == l {
//             ans[i] = l
//         }
//     }
//     return ans
// }

type LCA struct {
	g  map[int]map[int]int
	a  map[int][]int
	ti map[int]int
	to map[int]int
	tm int
	dp map[int]int
	ds map[int]int
}

func newLCA(graph map[int]map[int]int, rt int) *LCA {
	l := &LCA{
		g:  graph,
		a:  make(map[int][]int),
		ti: make(map[int]int),
		to: make(map[int]int),
		tm: 0,
		dp: make(map[int]int),
		ds: make(map[int]int),
	}
	l.dfs(rt, rt, []int{rt}, 0)
	return l
}

func (l *LCA) dfs(v, p int, pt []int, d int) {
	l.tm++
	l.ti[v] = l.tm
	up := 1
	for len(pt) >= up {
		l.a[v] = append(l.a[v], pt[len(pt)-up])
		up *= 2
	}
	l.dp[v] = len(pt)
	l.ds[v] = d
	pt = append(pt, v)
	if nb, ok := l.g[v]; ok {
		for u, w := range nb {
			if u != p {
				l.dfs(u, v, pt, d+w)
			}
		}
	}
	pt = pt[:len(pt)-1]
	l.tm++
	l.to[v] = l.tm
}

func (l *LCA) isAncestor(u, v int) bool {
	return l.ti[u] <= l.ti[v] && l.to[u] >= l.to[v]
}

func (l *LCA) lca(u, v int) int {
	if l.isAncestor(u, v) {
		return u
	}
	if l.isAncestor(v, u) {
		return v
	}
	d := len(l.a[u]) - 1
	for d >= 0 {
		if d >= len(l.a[u]) {
			d = len(l.a[u]) - 1
		}
		if d >= 0 && !l.isAncestor(l.a[u][d], v) {
			u = l.a[u][d]
		}
		d--
	}
	if len(l.a[u]) > 0 {
		return l.a[u][0]
	}
	return u
}

func (l *LCA) distance(u, v int) int {
	a := l.lca(u, v)
	if a == u || a == v {
		return int(math.Abs(float64(l.ds[v] - l.ds[u])))
	}
	return (l.ds[u] - l.ds[a]) + (l.ds[v] - l.ds[a])
}

func (l *LCA) findNodeDistFloat(u int, dt float64, md int) int {
	d := len(l.a[u]) - 1
	m := u
	for d >= 0 {
		if d >= len(l.a[u]) {
			d = len(l.a[u]) - 1
		}
		if d >= 0 {
			v := l.a[u][d]
			if float64(l.ds[v]) >= dt {
				m = v
				u = l.a[u][d]
			}
		}
		d--
	}
	if md == 0 || math.Abs(float64(l.ds[m])-dt) < 1e-9 {
		return m
	}
	if len(l.a[m]) > 0 {
		return l.a[m][0]
	}
	return m
}

func (l *LCA) median(u, v int) int {
	gl := float64(l.distance(u, v)) / 2.0
	a := l.lca(u, v)
	if u == a {
		return l.findNodeDistFloat(v, gl+float64(l.ds[u]), 0)
	} else if v == a {
		return l.findNodeDistFloat(u, gl+float64(l.ds[v]), 1)
	} else {
		d := l.distance(a, u)
		if float64(d) >= gl {
			return l.findNodeDistFloat(u, float64(d)-gl+float64(l.ds[a]), 1)
		} else {
			return l.findNodeDistFloat(v, gl-float64(d)+float64(l.ds[a]), 0)
		}
	}
}

func findMedian(n int, edges [][]int, queries [][]int) []int {
	g := make(map[int]map[int]int)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if g[u] == nil {
			g[u] = make(map[int]int)
		}
		if g[v] == nil {
			g[v] = make(map[int]int)
		}
		g[u][v] = w
		g[v][u] = w
	}
	l := newLCA(g, 0)
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = l.median(q[0], q[1])
	}
	return res
}

/*
This implementation uses a binary lifting approach for efficient Lowest Common Ancestor (LCA) queries combined
with weighted median computation on tree paths. The algorithm performs a depth-first search during initialization
to build ancestor tables where each node stores its 2^i-th ancestors, enabling logarithmic-time LCA queries
through binary jumping. The core logic determines the weighted median by calculating the total path weight between
two query nodes, finding their LCA, and then using binary search with the precomputed ancestor tables to locate
the exact node where cumulative weight from the starting point reaches or exceeds half the total path weight. The
time complexity is O(n log n) for preprocessing to build the binary lifting tables and O(log n) per query for LCA
computation and median finding, while the space complexity is O(n log n) due to storing log n ancestors for each
of the n nodes in the ancestor tables, plus O(n) for auxiliary data structures storing distances, depths, and
timing information.
*/