func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt
			}
		}
	}
	dp[0][0][2] = coins[0][0]
	if coins[0][0] < 0 {
		dp[0][0][1] = 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= 2; k++ {
				if dp[i][j][k] == math.MinInt {
					continue
				}
				if j+1 < n {
					dp[i][j+1][k] = max(dp[i][j+1][k], dp[i][j][k]+coins[i][j+1])
					if k > 0 && coins[i][j+1] < 0 {
						dp[i][j+1][k-1] = max(dp[i][j+1][k-1], dp[i][j][k])
					}
				}
				if i+1 < m {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k]+coins[i+1][j])
					if k > 0 && coins[i+1][j] < 0 {
						dp[i+1][j][k-1] = max(dp[i+1][j][k-1], dp[i][j][k])
					}
				}
			}
		}
	}
	res := math.MinInt
	for k := 0; k <= 2; k++ {
		res = max(res, dp[m-1][n-1][k])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}