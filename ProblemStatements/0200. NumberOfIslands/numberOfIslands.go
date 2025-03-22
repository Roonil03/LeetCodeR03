func numIslands(grid [][]byte) int {
	res := 0
	r, c := len(grid), len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				res++
				explore(grid, i, j, r, c)
			}
		}
	}
	return res
}

func explore(grid [][]byte, i, j, r, c int) {
	if i < 0 || i >= r || j < 0 || j >= c || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	explore(grid, i+1, j, r, c)
	explore(grid, i, j+1, r, c)
	explore(grid, i-1, j, r, c)
	explore(grid, i, j-1, r, c)
}