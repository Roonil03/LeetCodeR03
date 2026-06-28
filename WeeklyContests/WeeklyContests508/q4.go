func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	adj := make([][][2]int64, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int64{int64(e[1]), int64(e[2])})
	}
	pt := make([]int64, n)
	for i := range pt {
		pt[i] = int64(-1)
	}
	g1 := func(a, b xxx) bool {
		if a.t == b.t {
			return a.p > b.p
		}
		return a.t < b.t
	}
	pq := NewPriorityQueue[xxx](g1)
	pq.Push(xxx{0, int64(power), int64(source)})
	for pq.Len() > 0 {
		c := pq.Pop()
		if c.p <= pt[c.u] {
			continue
		}
		pt[c.u] = c.p
		if c.u == int64(target) {
			return []int64{int64(c.t), int64(c.p)}
		}
		if c.p >= int64(cost[c.u]) {
			p := c.p - int64(cost[c.u])
			for _, e := range adj[c.u] {
				v, w := e[0], e[1]
				if p > pt[v] {
					pq.Push(xxx{c.t + w, p, v})
				}
			}
		}
	}
	return []int64{-1, -1}
}

type xxx struct {
	t, p, u int64
}

type xx xxx

func (a xx) Less(b xx) bool {
	if a.t == b.t {
		return a.p > b.p
	}
	return a.t < b.t
}

type PriorityQueue[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0), less: less}
}

func (h *PriorityQueue[T]) Push(x T) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *PriorityQueue[T]) Pop() T {
	n := len(h.data) - 1
	h.data[0], h.data[n] = h.data[n], h.data[0]
	x := h.data[n]
	h.data = h.data[0:n]
	h.down(0, n)
	return x
}

func (h *PriorityQueue[T]) Len() int {
	return len(h.data)
}

func (h *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func (h *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.less(h.data[j2], h.data[j1]) {
			j = j2
		}
		if !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
	return i > i0
}