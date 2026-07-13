func maxConsistentColumns(grid [][]int, limit int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	res := 1
	for i := range n {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if dp[j]+1 <= dp[i] {
				continue
			}
			fg := true
			for k := range m {
				if a := grid[k][i] - grid[k][j]; a < -limit || a > limit {
					fg = false
					break
				}
			}
			if fg {
				dp[i] = dp[j] + 1
			}
		}
		res = max(res, dp[i])
	}
	return res
}