func findMaxFish(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	maxFish := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				maxFish = max(maxFish, dfs(i, j, grid, n, m))
			}
		}
	}
	return maxFish
}

var d = []int{0, 1, 0, -1, 0}

func dfs(i, j int, grid [][]int, n, m int) int {
	f := grid[i][j]
	grid[i][j] = 0
	for k := 0; k < 4; k++ {
		x, y := i+d[k], j+d[k+1]
		if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] > 0 {
			f += dfs(x, y, grid, n, m)
		}
	}
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
