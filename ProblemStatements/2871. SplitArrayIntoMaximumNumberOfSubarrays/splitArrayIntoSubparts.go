func maxSubarrays(nums []int) int {
	res, cur := 0, math.MaxInt32
	for _, n := range nums {
		cur &= n
		if cur == 0 {
			res++
			cur = math.MaxInt32
		}
	}
	if res == 0 {
		return 1
	}
	return res
}