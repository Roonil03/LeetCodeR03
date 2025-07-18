func minimumDifference(nums []int) int64 {
	a, b := 0, 0
	n := len(nums)
	k := n / 3
	res := int64(1 << 62)
	h1 := &maxHeap{}
	h2 := &minHeap{}
	temp := make([]int, n)
	for i := n - 1; i >= k; i-- {
		heap.Push(h2, nums[i])
		b += nums[i]
		if h2.Len() > k {
			b -= heap.Pop(h2).(int)
		}
		if h2.Len() == k {
			temp[i] = b
		}
	}
	for i := 0; i < n-k; i++ {
		heap.Push(h1, nums[i])
		a += nums[i]
		if h1.Len() > k {
			a -= heap.Pop(h1).(int)
		}
		if h1.Len() == k {
			if i+1 < n {
				res = min(res, int64(a-temp[i+1]))
			}
		}
	}
	return int64(res)
}

type maxHeap []int

func (maxH maxHeap) Len() int           { return len(maxH) }
func (maxH maxHeap) Less(i, j int) bool { return maxH[i] > maxH[j] }
func (maxH maxHeap) Swap(i, j int)      { maxH[i], maxH[j] = maxH[j], maxH[i] }
func (maxH *maxHeap) Push(x any)        { *maxH = append(*maxH, x.(int)) }
func (maxH *maxHeap) Pop() any {
	old := *maxH
	n := len(old)
	x := old[n-1]
	*maxH = old[0 : n-1]
	return x
}

type minHeap []int

func (minH minHeap) Len() int           { return len(minH) }
func (minH minHeap) Less(i, j int) bool { return minH[i] < minH[j] }
func (minH minHeap) Swap(i, j int)      { minH[i], minH[j] = minH[j], minH[i] }
func (minH *minHeap) Push(x any)        { *minH = append(*minH, x.(int)) }
func (minH *minHeap) Pop() any {
	old := *minH
	n := len(old)
	x := old[n-1]
	*minH = old[0 : n-1]
	return x
}