func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for _, n := range nums {
		sum += n
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}