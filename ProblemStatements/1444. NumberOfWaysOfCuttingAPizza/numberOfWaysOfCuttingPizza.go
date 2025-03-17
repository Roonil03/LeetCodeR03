func ways(pizza []string, k int) int {
	rows, cols := len(pizza), len(pizza[0])
	mod := 1000000007
	apples := make([][]int, rows+1)
	for i := range apples {
		apples[i] = make([]int, cols+1)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			apples[i][j] = apples[i+1][j] + apples[i][j+1] - apples[i+1][j+1]
			if pizza[i][j] == 'A' {
				apples[i][j]++
			}
		}
	}
	dp := make([][][]int, k)
	for i := range dp {
		dp[i] = make([][]int, rows)
		for j := range dp[i] {
			dp[i][j] = make([]int, cols)
			for c := range dp[i][j] {
				dp[i][j][c] = -1
			}
		}
	}
	var dfs func(pieces, r, c int) int
	dfs = func(pieces, r, c int) int {
		if apples[r][c] == 0 {
			return 0
		}
		if pieces == 1 {
			return 1
		}
		if dp[pieces-1][r][c] != -1 {
			return dp[pieces-1][r][c]
		}
		ans := 0
		for nr := r + 1; nr < rows; nr++ {
			if apples[r][c]-apples[nr][c] > 0 {
				ans = (ans + dfs(pieces-1, nr, c)) % mod
			}
		}
		for nc := c + 1; nc < cols; nc++ {
			if apples[r][c]-apples[r][nc] > 0 {
				ans = (ans + dfs(pieces-1, r, nc)) % mod
			}
		}
		dp[pieces-1][r][c] = ans
		return ans
	}
	return dfs(k, 0, 0)
}