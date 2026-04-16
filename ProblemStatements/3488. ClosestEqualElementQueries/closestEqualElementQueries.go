func solveQueries(nums []int, queries []int) []int {
	lm, rm := make(map[int]int), make(map[int]int)
	dist, res := make([]int, len(nums)), make([]int, 0, len(queries))
	n := len(nums)
	for i := range nums {
		dist[i] = math.MaxInt
	}
	for i := 0; i < n*2; i++ {
		id := i % n
		val := nums[id]
		if prev, ok := lm[val]; ok {
			dist[id] = i - prev
		}
		lm[val] = i
	}
	for i := n*2 - 1; i >= 0; i-- {
		id := i % n
		val := nums[id]
		if next, ok := rm[val]; ok {
			if next-i < dist[id] {
				dist[id] = next - i
			}
		}
		rm[val] = i
	}
	for _, q := range queries {
		if dist[q] >= n {
			res = append(res, -1)
		} else {
			res = append(res, dist[q])
		}
	}
	return res
}