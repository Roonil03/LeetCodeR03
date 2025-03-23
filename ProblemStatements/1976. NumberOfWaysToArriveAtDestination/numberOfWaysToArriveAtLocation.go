type Edge struct {
	to, time int
}

type State struct {
	node, d int
}

type pq []State

func (pq pq) Len() int {
	return len(pq)
}

func (pq pq) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}

func (pq pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *pq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func countPaths(n int, roads [][]int) int {
	const mod = 1e9 + 7
	graph := make([][]Edge, n)
	for _, road := range roads {
		u, v, time := road[0], road[1], road[2]
		graph[u] = append(graph[u], Edge{v, time})
		graph[v] = append(graph[v], Edge{u, time})
	}
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = 1<<63 - 1
	}
	dist[0] = 0
	ways := make([]int64, n)
	ways[0] = 1
	pq := &pq{{0, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(State)
		node, d := curr.node, curr.d
		if int64(d) > dist[node] {
			continue
		}
		for _, edge := range graph[node] {
			newDist := dist[node] + int64(edge.time)
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				ways[edge.to] = ways[node]
				heap.Push(pq, State{edge.to, int(newDist)})
			} else if newDist == dist[edge.to] {
				ways[edge.to] = (ways[edge.to] + ways[node]) % mod
			}
		}
	}
	return int(ways[n-1])
}