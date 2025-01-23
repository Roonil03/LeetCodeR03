func countServers(grid [][]int) int {
	r := make([]int, len(grid))
	c := make([]int, len(grid[0]))
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				r[i]++
				c[j]++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (r[i] > 1 || c[j] > 1) {
				res++
			}
		}
	}
	return res
}