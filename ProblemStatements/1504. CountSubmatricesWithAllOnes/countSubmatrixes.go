func numSubmat(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
			} else {
				dp[i][j] = 0
			}
			a := dp[i][j]
			for k := i; k >= 0 && a > 0; k-- {
				if dp[k][j] < a {
					a = dp[k][j]
				}
				res += a
			}
		}
	}
	return res
}