func leftRightDifference(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	r, l := 0, 0
	for i := 0; i < n; i++ {
		r += nums[i]
	}
	for i := 0; i < n; i++ {
		r -= nums[i]
		res[i] = abs(l - r)
		l += nums[i]
	}
	return res
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}