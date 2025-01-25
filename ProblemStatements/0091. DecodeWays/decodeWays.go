func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	for i := 2; i <= n; i++ {
		a, err1 := strconv.Atoi(s[i-1 : i])
		b, err2 := strconv.Atoi(s[i-2 : i])
		if err1 != nil && err2 != nil {
			return dp[0]
		}
		if a >= 1 && a <= 9 {
			dp[i] += dp[i-1]
		}
		if b >= 10 && b <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}