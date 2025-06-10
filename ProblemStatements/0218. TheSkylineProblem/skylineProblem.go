func getSkyline(buildings [][]int) [][]int {
	e := make([][]int, 0, len(buildings)*2)
	for _, b := range buildings {
		e = append(e, []int{b[0], -b[2]})
		e = append(e, []int{b[1], b[2]})
	}
	sort.Slice(e, func(i, j int) bool {
		if e[i][0] != e[j][0] {
			return e[i][0] < e[j][0]
		}
		return e[i][1] < e[j][1]
	})
	res := make([][]int, 0)
	h := &MaxHeap{}
	heap.Init(h)
	heap.Push(h, []int{0})
	prevHeight := 0
	for _, event := range e {
		x, height := event[0], event[1]
		if height < 0 {
			heap.Push(h, []int{-height})
		} else {
			for i := 0; i < h.Len(); i++ {
				if (*h)[i][0] == height {
					heap.Remove(h, i)
					break
				}
			}
		}
		currHeight := (*h)[0][0]
		if currHeight != prevHeight {
			res = append(res, []int{x, currHeight})
			prevHeight = currHeight
		}
	}
	return res
}

type MaxHeap [][]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
