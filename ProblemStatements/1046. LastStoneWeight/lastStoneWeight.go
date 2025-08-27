func lastStoneWeight(stones []int) int {
	m := NewMaxHeap()
	for _, i := range stones {
		heap.Push(m, i)
	}
	for m.Len() > 1 {
		s1 := heap.Pop(m).(int)
		s2 := heap.Pop(m).(int)
		if s1 != s2 {
			heap.Push(m, s1-s2)
		}
	}
	res := 0
	if m.Len() == 1 {
		res = (*m)[0]
	}
	return res
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewMaxHeap() *MaxHeap {
	h := &MaxHeap{}
	heap.Init(h)
	return h
}