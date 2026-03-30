func minMoves(nums []int) int {
	m := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		m = min(v, m)
	}
	return sum - m*len(nums)
}