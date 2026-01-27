func minCost(n int, edges [][]int) int {
	fwd := make([][][2]int, n)
	inc := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		fwd[u] = append(fwd[u], [2]int{v, w})
		inc[v] = append(inc[v], [2]int{u, w})
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[0] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Point{u: 0, w: 0})
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Point)
		u := cur.u
		w := cur.w
		// i := 0
		// if used{
		//     i = 1
		// }
		if w > dist[u] {
			continue
		}
		for _, e := range fwd[u] {
			nbr := e[0]
			w1 := e[1]
			cost := w + w1
			if cost < dist[nbr] {
				dist[nbr] = cost
				heap.Push(&pq, &Point{u: nbr, w: cost})
				fmt.Println("u: ", u, " v: ", nbr, " cost: ", cost)
			}
		}
		// if !used{
		for _, e := range inc[u] {
			nbr := e[0]
			w1 := e[1]
			cost := w + 2*w1
			if cost < dist[nbr] {
				dist[nbr] = cost
				heap.Push(&pq, &Point{u: nbr, w: cost /*, used: true*/})
				fmt.Println("u: ", u, " v: ", nbr, " cost: ", cost)
			}
		}
		// }
	}
	// res := min(dist[n-1], dist[n-1])
	if dist[n-1] == math.MaxInt {
		return -1
	}
	return dist[n-1]
}

type Point struct {
	u  int
	w  int
	id int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w < pq[j].w
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].id = i
	pq[j].id = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	state := x.(*Point)
	state.id = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil
	state.id = -1
	*pq = old[0 : n-1]
	return state
}