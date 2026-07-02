import "cmp"

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dist := make([]int, m*n)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}
	a := grid[0][0]
	if a >= health {
		return false
	}
	dist[0] = a
	pq := NewPriorityQueue[int64]()
	pq.Push(int64(a) << 32)
	dirs := []int{-1, 0, 1, 0, -1}
	for pq.Len() > 0 {
		cur := pq.Pop()
		cost := int(cur >> 32)
		r := int((cur >> 16) & 0xFFFF)
		c := int(cur & 0xFFFF)
		if r == m-1 && c == n-1 {
			return true
		}
		if cost > dist[r*n+c] {
			continue
		}
		for i := range 4 {
			nr, nc := r+dirs[i], c+dirs[i+1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				ncost := cost + grid[nr][nc]
				if ncost < dist[nr*n+nc] && ncost < health {
					dist[nr*n+nc] = ncost
					pq.Push((int64(ncost) << 32) | (int64(nr) << 16) | int64(nc))
				}
			}
		}
	}
	return false
}

type PriorityQueue[T cmp.Ordered] struct {
	data []T
}

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0)}
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
		if i == j || !(h.data[j] < h.data[i]) {
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
		if j2 := j1 + 1; j2 < n && h.data[j2] < h.data[j1] {
			j = j2
		}
		if !(h.data[j] < h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
	return i > i0
}