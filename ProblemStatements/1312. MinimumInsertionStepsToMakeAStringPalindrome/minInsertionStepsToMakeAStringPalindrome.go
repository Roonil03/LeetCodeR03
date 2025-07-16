func minInsertions(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j+i-1 < n; j++ {
			k := j + i - 1
			if s[j] == s[k] {
				if i == 2 {
					dp[j][k] = 2
				} else {
					dp[j][k] = dp[j+1][k-1] + 2
				}
			} else {
				if dp[j+1][k] > dp[j][k-1] {
					dp[j][k] = dp[j+1][k]
				} else {
					dp[j][k] = dp[j][k-1]
				}
			}
		}
	}
	return n - dp[0][n-1]
}