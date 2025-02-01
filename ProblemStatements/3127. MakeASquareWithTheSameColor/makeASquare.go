func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if (j-1 >= 0 && grid[i][j] == grid[i][j-1]) || (j+1 < 3 && grid[i][j] == grid[i][j+1]) {
				if (i+1 < 3 && grid[i][j] == grid[i+1][j]) || (i-1 >= 0 && grid[i][j] == grid[i-1][j]) {
					return true
				}
			}
		}
	}
	return false
}