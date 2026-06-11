func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	w := bits.Len(uint(n))
	Min := make([][]int, w)
	Max := make([][]int, w)
	for i := 0; i < w; i++ {
		Min[i] = make([]int, n)
		Max[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		Min[0][i] = nums[i]
		Max[0][i] = nums[i]
	}
	for i := 1; i < w; i++ {
		for j := 0; j+(1<<i) <= n; j++ {
			Min[i][j] = min(Min[i-1][j], Min[i-1][j+(1<<(i-1))])
			Max[i][j] = max(Max[i-1][j], Max[i-1][j+(1<<(i-1))])
		}
	}
	query := func(left, right int) int {
		if left >= right {
			return 0
		}
		kBits := bits.Len(uint(right-left)) - 1
		mx := max(Max[kBits][left], Max[kBits][right-(1<<kBits)])
		mn := min(Min[kBits][left], Min[kBits][right-(1<<kBits)])
		return mx - mn
	}
	hp := &MaxHeap{}
	heap.Init(hp)
	for i := 0; i < n; i++ {
		heap.Push(hp, gg{l: i, r: n, tot: query(i, n)})
	}
	var res int64
	c := 0
	for hp.Len() > 0 && c < k {
		top := heap.Pop(hp).(gg)
		if top.tot == 0 {
			break
		}
		res += int64(top.tot)
		c++
		heap.Push(hp, gg{l: top.l, r: top.r - 1, tot: query(top.l, top.r-1)})
	}
	return res
}

type gg struct {
	l, r int
	m, n int
	tot  int
}

type MaxHeap []gg

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].tot == h[j].tot {
		if h[i].l == h[j].l {
			return h[i].r > h[j].r
		}
		return h[i].l > h[j].l
	}
	return h[i].tot > h[j].tot
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(gg))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}