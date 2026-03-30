func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([]int, n+m)
	for i := m - 1; i >= 0; i-- {
		t := make([]int, m+1)
		for l := i; l >= 0; l-- {
			r := n - 1 - (i - l)
			a := multipliers[i]*nums[l] + dp[l+1]
			b := multipliers[i]*nums[r] + dp[l]
			t[l] = max(a, b)
		}
		dp = t
	}
	return dp[0]
}