// func minOperations(nums []int, k int) int {
//     h := &IntHeap{}
//     heap.Init(h)
//     for _, num := range nums{
//         heap.Push(h, num)
//     }
//     a := 0
//     for h.Len() > 0 && (*h)[0] < k {
//         if h.Len() < 2{
//             return -1
//         }
//         x := heap.Pop(h).(int)
//         y := heap.Pop(h).(int)
//         z := 2*x + y
//         heap.Push(h, z)
//         a++
//     }
//     return a
// }

func minOperations(nums []int, k int) int {
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for _, num := range nums {
		if num < k {
			heap.Push(minHeap, num)
		}
	}
	opsCount := 0
	for minHeap.Len() > 0 {
		first := heap.Pop(minHeap).(int)
		opsCount++
		if minHeap.Len() == 0 {
			break
		}
		second := heap.Pop(minHeap).(int)
		newNumber := 2*first + second
		if newNumber < k {
			heap.Push(minHeap, newNumber)
		}
	}
	return opsCount
}

//Heap:
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}