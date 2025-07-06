type State struct {
	row, col int
	odd      bool
	cost     int64
}

func minCost(m int, n int, waitCost [][]int) int64 {
	dirs := [][]int{{0, 1}, {1, 0}}
	d := make([][][]int64, m)
	for i := range d {
		d[i] = make([][]int64, n)
		for j := range d[i] {
			d[i][j] = make([]int64, 2)
			d[i][j][0] = math.MaxInt64
			d[i][j][1] = math.MaxInt64
		}
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	d[0][0][1] = 1
	heap.Push(pq, &State{0, 0, true, 1})
	for pq.Len() > 0 {
		c := heap.Pop(pq).(*State)
		id := 0
		if c.odd {
			id = 1
		}
		if c.cost > d[c.row][c.col][id] {
			continue
		}
		if c.odd {
			for _, d1 := range dirs {
				nr, nc := c.row+d1[0], c.col+d1[1]
				if nr < m && nc < n {
					e := (nr + 1) * (nc + 1)
					ct := c.cost + int64(e)
					if ct < d[nr][nc][0] {
						d[nr][nc][0] = ct
						heap.Push(pq, &State{nr, nc, false, ct})
					}
				}
			}
		} else {
			ct := waitCost[c.row][c.col]
			a := c.cost + int64(ct)
			if a < d[c.row][c.col][1] {
				d[c.row][c.col][1] = a
				heap.Push(pq, &State{c.row, c.col, true, a})
			}
		}
	}
	return min64(d[m-1][n-1][0], d[m-1][n-1][1])
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}