func networkDelayTime(times [][]int, n int, k int) int {
    adj := make([][][2]int, n+1)
    for _, t := range times {
        u, v, w := t[0], t[1], t[2]
        adj[u] = append(adj[u], [2]int{v, w})
    }
    dist := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dist[i] = math.MaxInt32
    }
    dist[k] = 0
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &Item{k, 0})
    for pq.Len() > 0 {
        cur := heap.Pop(pq).(*Item)
        u, d := cur.node, cur.dist
        
        if d > dist[u] {
            continue
        }
        
        for _, e := range adj[u] {
            v, w := e[0], e[1]
            if dist[u]+w < dist[v] {
                dist[v] = dist[u] + w
                heap.Push(pq, &Item{v, dist[v]})
            }
        }
    }
    res := 0
    for i := 1; i <= n; i++ {
        if dist[i] == 1<<31 - 1 {
            return -1
        }
        if dist[i] > res {
            res = dist[i]
        }
    }    
    return res
}

type Item struct {
    node int
    dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}func networkDelayTime(times [][]int, n int, k int) int {
    adj := make([][][2]int, n+1)
    for _, t := range times {
        u, v, w := t[0], t[1], t[2]
        adj[u] = append(adj[u], [2]int{v, w})
    }
    dist := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dist[i] = math.MaxInt32
    }
    dist[k] = 0
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &Item{k, 0})
    for pq.Len() > 0 {
        cur := heap.Pop(pq).(*Item)
        u, d := cur.node, cur.dist
        
        if d > dist[u] {
            continue
        }
        
        for _, e := range adj[u] {
            v, w := e[0], e[1]
            if dist[u]+w < dist[v] {
                dist[v] = dist[u] + w
                heap.Push(pq, &Item{v, dist[v]})
            }
        }
    }
    res := 0
    for i := 1; i <= n; i++ {
        if dist[i] == 1<<31 - 1 {
            return -1
        }
        if dist[i] > res {
            res = dist[i]
        }
    }    
    return res
}

type Item struct {
    node int
    dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

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