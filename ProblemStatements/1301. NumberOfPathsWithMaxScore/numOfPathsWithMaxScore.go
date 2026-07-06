func pathsWithMaxScore(board []string) []int {
	n := len(board)
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	ct := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		ct[i] = make([]int, n+1)
	}
	ct[n-1][n-1] = 1
	dirs := [3][2]int{{0, 1}, {1, 0}, {1, 1}}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' || (i == n-1 && j == n-1) {
				continue
			}
			mx := -1
			for _, v := range dirs {
				ni, nj := i+v[0], j+v[1]
				if ct[ni][nj] > 0 && dp[ni][nj] > mx {
					mx = dp[ni][nj]
				}
			}
			if mx != -1 {
				for _, v := range dirs {
					ni, nj := i+v[0], j+v[1]
					if ct[ni][nj] > 0 && dp[ni][nj] == mx {
						ct[i][j] = (ct[i][j] + ct[ni][nj]) % mod
					}
				}
				if board[i][j] != 'E' {
					dp[i][j] = mx + int(board[i][j]-'0')
				} else {
					dp[i][j] = mx
				}
			}
		}
	}
	if ct[0][0] == 0 {
		return []int{0, 0}
	}
	return []int{dp[0][0], ct[0][0]}
}