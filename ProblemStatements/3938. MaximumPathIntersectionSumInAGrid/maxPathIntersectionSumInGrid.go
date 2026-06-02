func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := grid[0][0] + grid[0][1]
	for i := range m {
		dp := grid[i][0]
		for j := 1; j < n; j++ {
			v := dp + grid[i][j]
			if v > res {
				res = v
			}
			if grid[i][j] > v {
				dp = grid[i][j]
			} else {
				dp = v
			}
		}
		if i > 0 && i < m-1 {
			for j := 1; j < n-1; j++ {
				if grid[i][j] > res {
					res = grid[i][j]
				}
			}
		}
	}
	for j := range n {
		dp := grid[0][j]
		for i := 1; i < m; i++ {
			v := dp + grid[i][j]
			if v > res {
				res = v
			}
			if grid[i][j] > v {
				dp = grid[i][j]
			} else {
				dp = v
			}
		}
		if j > 0 && j < n-1 {
			for i := 1; i < m-1; i++ {
				if grid[i][j] > res {
					res = grid[i][j]
				}
			}
		}
	}
	return res
}