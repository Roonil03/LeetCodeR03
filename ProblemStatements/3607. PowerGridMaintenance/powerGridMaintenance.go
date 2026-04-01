func processQueries(c int, connections [][]int, queries [][]int) []int {
    dsu := NewDSU(c)
    for _, i := range connections{
        u, v := i[0] - 1, i[1] - 1
        dsu.Union(u, v)
    }
    grid := make([][]int, c)
    for i := range c{
        rt := dsu.Find(i)
        grid[rt] = append(grid[rt], i + 1)
    }
    ptr := make([]int, c)
    on := make([]bool, c +1)
    for i := 1; i <= c; i++{
        on[i] = true
    }
    res := make([]int, 0)
    for _, q := range queries{
        t, x := q[0], q[1]
        if t == 2{
            on[x] = false
            continue
        }
        if on[x]{
            res = append(res, x)
            continue
        }
        rt := dsu.Find(x - 1)
        n := grid[rt]
        for ptr[rt] < len(n) && !on[n[ptr[rt]]]{
            ptr[rt]++
        }
        if ptr[rt] == len(n){
            res = append(res, -1)
        } else{
            res = append(res, n[ptr[rt]])
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