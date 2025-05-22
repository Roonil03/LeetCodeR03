// func maxRemoval(nums []int, queries [][]int) int {
//     n := len(nums)
// 	m := len(queries)
// 	left, right := 0, m
// 	ans := -1
// 	isPossible := func(removed map[int]bool) bool {
// 		diff := make([]int, n+2)
// 		for i := 0; i < n; i++ {
// 			diff[i+1] = nums[i]
// 			diff[i+1] -= diff[i]
// 		}
// 		for i := 0; i < m; i++ {
// 			if removed[i] {
// 				continue
// 			}
// 			l, r := queries[i][0]+1, queries[i][1]+1
// 			diff[l] -= 1
// 			diff[r+1] += 1
// 		}
// 		total := 0
// 		for i := 1; i <= n; i++ {
// 			total += diff[i]
// 			if total != 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	for left <= right {
// 		mid := (left + right) / 2
// 		found := false
// 		indices := make([]int, m)
// 		for i := 0; i < m; i++ {
// 			indices[i] = i
// 		}
// 		var dfs func(pos int, cnt int, removed map[int]bool)
// 		dfs = func(pos int, cnt int, removed map[int]bool) {
// 			if found {
// 				return
// 			}
// 			if cnt == mid {
// 				if isPossible(removed) {
// 					found = true
// 				}
// 				return
// 			}
// 			if pos == m {
// 				return
// 			}
// 			removed[indices[pos]] = true
// 			dfs(pos+1, cnt+1, removed)
// 			delete(removed, indices[pos])
// 			dfs(pos+1, cnt, removed)
// 		}
// 		removed := map[int]bool{}
// 		dfs(0, 0, removed)
// 		if found {
// 			ans = mid
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return ans
// }

// type MaxHeap []int
// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	x := old[len(old)-1]
// 	*h = old[:len(old)-1]
// 	return x
// }

// type MinHeap []int
// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	x := old[len(old)-1]
// 	*h = old[:len(old)-1]
// 	return x
// }

// func maxRemoval(nums []int, queries [][]int) int {
// 	sort.Slice(queries, func(i, j int) bool {
// 		if queries[i][0] == queries[j][0] {
// 			return queries[i][1] < queries[j][1]
// 		}
// 		return queries[i][0] < queries[j][0]
// 	})
// 	n := len(nums)
// 	candidate := &MaxHeap{}
// 	heap.Init(candidate)
// 	chosen := &MinHeap{}
// 	heap.Init(chosen)
// 	ans, j := 0, 0
// 	for i := 0; i < n; i++ {
// 		for j < len(queries) && queries[j][0] == i {
// 			heap.Push(candidate, queries[j][1])
// 			j++
// 		}
// 		nums[i] -= chosen.Len()
// 		for nums[i] > 0 && candidate.Len() > 0 && (*candidate)[0] >= i {
// 			ans++
// 			heap.Push(chosen, heap.Pop(candidate).(int))
// 			nums[i]--
// 		}
// 		if nums[i] > 0 {
// 			return -1
// 		}
// 		for chosen.Len() > 0 && (*chosen)[0] == i {
// 			heap.Pop(chosen)
// 		}
// 	}
// 	return len(queries) - ans
// }

func maxRemoval(nums []int, queries [][]int) int {
	slices.SortFunc(queries, func(a, b []int) int { return a[0] - b[0] })
	h := hp{}
	diff := make([]int, len(nums)+1)
	sumD, j := 0, 0
	for i, x := range nums {
		sumD += diff[i]
		for ; j < len(queries) && queries[j][0] <= i; j++ {
			heap.Push(&h, queries[j][1])
		}
		for sumD < x && h.Len() > 0 && h.IntSlice[0] >= i {
			sumD++
			diff[heap.Pop(&h).(int)+1]--
		}
		if sumD < x {
			return -1
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }