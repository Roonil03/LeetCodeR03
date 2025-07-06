type Event struct {
	to, start, end int
}

type State struct {
	node, time int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func minTime(n int, edges [][]int) int {
	adj := make([][]Event, n)
	// tem := 0
	for _, e := range edges {
		u, v, s, t := e[0], e[1], e[2], e[3]
		adj[u] = append(adj[u], Event{v, s, t})
		// if t > tem{
		//     tem = t
		// }
	}
	// t := tem + n + 5
	// vis := make([][]bool, n)
	// for i:= 0; i < n; i++{
	//     vis[i] = make([]bool, t)
	// }
	// pq := &PQ{}
	// heap.Init(pq)
	// heap.Push(pq, State{0,0})
	// vis[0][0] = true
	// for pq.Len() > 0{
	//     c := heap.Pop(pq).(State)
	//     u, v := c.node, c.time
	//     if u == n - 1{
	//         return v
	//     }
	//     if v + 1 < t && !vis[u][v+1]{
	//         vis[u][v + 1] = true
	//         heap.Push(pq, State{u, v + 1})
	//     }
	//     for _, e := range adj[u]{
	//         if v >= e.start && v <= e.end && v + 1 < t && !vis[e.to][v+1]{
	//             vis[e.to][v+1] = true
	//             heap.Push(pq, State{e.to, v + 1})
	//         }
	//     }
	// }
	tem := make([]int, n)
	for i := range tem {
		tem[i] = math.MaxInt64
	}
	tem[0] = 0
	pq := &PriorityQueue{{0, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		c := heap.Pop(pq).(State)
		u, t := c.node, c.time
		if u == n-1 {
			return t
		}
		if t > tem[u] {
			continue
		}
		for _, e := range adj[u] {
			dep := t
			if dep < e.start {
				dep = e.start
			}
			if dep <= e.end {
				a := dep + 1
				if a < tem[e.to] {
					tem[e.to] = a
					heap.Push(pq, State{e.to, a})
				}
			}
		}
	}
	return -1
}