type Event [][]int

func (e Event) Len() int { return len(e) }
func (e Event) Less(i, j int) bool {
	if e[i][0] == e[j][0] {
		return e[i][1] < e[j][1]
	}
	return e[i][0] < e[j][0]
}
func (e Event) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Push(val int) {
	h.arr = append(h.arr, val)
	i := len(h.arr) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.arr[p] <= h.arr[i] {
			break
		}
		h.arr[p], h.arr[i] = h.arr[i], h.arr[p]
		i = p
	}
}

func (h *MinHeap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}
	res := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapify(0)
	return res
}

func (h *MinHeap) Peek() int {
	if len(h.arr) == 0 {
		return -1
	}
	return h.arr[0]
}

func (h *MinHeap) Len() int {
	return len(h.arr)
}

func (h *MinHeap) heapify(i int) {
	n := len(h.arr)
	for {
		smallest := i
		l := 2*i + 1
		r := 2*i + 2
		if l < n && h.arr[l] < h.arr[smallest] {
			smallest = l
		}
		if r < n && h.arr[r] < h.arr[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}

func maxEvents(events [][]int) int {
	sort.Sort(Event(events))
	a := 0
	for _, e := range events {
		if e[1] > a {
			a = e[1]
		}
	}
	h := &MinHeap{}
	res := 0
	d := 1
	id := 0
	n := len(events)
	for d <= a {
		for id < n && events[id][0] == d {
			h.Push(events[id][1])
			id++
		}
		for h.Len() > 0 && h.Peek() < d {
			h.Pop()
		}
		if h.Len() > 0 {
			h.Pop()
			res++
		}
		d++
	}
	return res
}
