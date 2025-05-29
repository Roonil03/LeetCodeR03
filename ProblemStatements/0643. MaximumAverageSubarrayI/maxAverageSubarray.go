func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	s := sum
	for i := k; i < n; i++ {
		sum += nums[i] - nums[i-k]
		if sum > s {
			s = sum
		}
	}
	return float64(s) / float64(k)
}