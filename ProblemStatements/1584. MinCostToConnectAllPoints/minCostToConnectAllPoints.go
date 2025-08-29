func minCostConnectPoints(points [][]int) int {
	n := len(points)
	vis := make([]bool, n)
	mc := make([]int, n)
	for i := 1; i < n; i++ {
		mc[i] = math.MaxInt32
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: 0, cost: 0})
	tot, ed := 0, 0
	for ed < n {
		item := heap.Pop(pq).(*Item)
		point := item.point
		cost := item.cost
		if vis[point] {
			continue
		}
		vis[point] = true
		tot += cost
		ed++
		for i := 0; i < n; i++ {
			if !vis[i] {
				dist := abs(points[point][0]-points[i][0]) + abs(points[point][1]-points[i][1])
				if dist < mc[i] {
					mc[i] = dist
					heap.Push(pq, &Item{point: i, cost: dist})
				}
			}
		}
	}
	return tot
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Item struct {
	point int
	cost  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
