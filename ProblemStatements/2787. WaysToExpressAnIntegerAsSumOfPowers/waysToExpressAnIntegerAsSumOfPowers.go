func numberOfWays(n int, x int) int {
	mod := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; ; i++ {
		p := 1
		for j := 0; j < x; j++ {
			p *= i
			if p > n {
				return dp[n]
			}
		}
		for j := n; j >= p; j-- {
			dp[j] = (dp[j] + dp[j-p]) % mod
		}
	}
}