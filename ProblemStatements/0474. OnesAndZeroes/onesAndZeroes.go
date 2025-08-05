func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		a, b := 0, 0
		for _, c := range s {
			if c == '0' {
				a++
			} else {
				b++
			}
		}
		for i := m; i >= a; i-- {
			for j := n; j >= b; j-- {
				if dp[i-a][j-b]+1 > dp[i][j] {
					dp[i][j] = dp[i-a][j-b] + 1
				}
			}
		}
	}
	return dp[m][n]
}