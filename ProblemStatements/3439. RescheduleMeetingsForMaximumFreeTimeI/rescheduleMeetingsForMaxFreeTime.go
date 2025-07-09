//type MinHeap []int

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	// n := len(startTime)
	// a := make([]int, n + 1)
	// for i := 1; i <= n; i++{
	//     a[i] = a[i-1] + (endTime[i-1] - startTime[i-1])
	// }
	// res := 0
	// // h := &MinHeap{};
	// // heap.Init(h)
	// // for r := 0; r <= n; r++{
	// //     a1 := a[r]
	// //     sum += a1
	// //     heap.Push(h, a1)
	// //     if h.Len() > k{
	// //         s := heap.Pop(h).(int)
	// //         sum -= s
	// //     }
	// //     if sum > res{
	// //         res = sum
	// //     }
	// // }
	// // // if res != 0{
	// // //     return res + 1
	// // // }
	// // return res
	// for i:= k - 1; i < n; i++{
	//     x, y := a[i+1] - a[i + 1 - k], eventTime
	//     if i != n - 1{
	//         y = startTime[i + 1]
	//     }
	//     z := 0
	//     if i != k - 1{
	//         z = endTime[i - k]
	//     }
	//     temp := y - z - x
	//     if temp > res{
	//         res = temp
	//     }
	// }
	// return res

	n := len(startTime)
	res := 0
	t := 0
	for i := 0; i < n; i++ {
		t += endTime[i] - startTime[i]
		var left int
		if i <= k-1 {
			left = 0
		} else {
			left = endTime[i-k]
		}
		var right int
		if i == n-1 {
			right = eventTime
		} else {
			right = startTime[i+1]
		}
		if right-left-t > res {
			res = right - left - t
		}
		if i >= k-1 {
			t -= endTime[i-k+1] - startTime[i-k+1]
		}
	}
	return res
}

// func (h MinHeap) Len() int            { return len(h) }
// func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
// func (h *MinHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[:n-1]
//     return x
// }