func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := range i {
			if dp[j] != -1 {
				dif := nums[i] - nums[j]
				if dif >= -target && dif <= target {
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[n-1]
}