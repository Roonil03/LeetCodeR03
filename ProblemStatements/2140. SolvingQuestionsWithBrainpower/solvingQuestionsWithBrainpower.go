func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		a := min(i+questions[i][1]+1, n)
		b := int64(questions[i][0]) + dp[a]
		dp[i] = max(dp[i+1], b)
	}
	return dp[0]
}