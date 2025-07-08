func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		c[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					c[i] = c[j]
				} else if dp[j]+1 == dp[i] {
					c[i] += c[j]
				}
			}
		}
	}
	ml := 0
	for i := 0; i < n; i++ {
		if dp[i] > ml {
			ml = dp[i]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == ml {
			res += c[i]
		}
	}
	return res
}