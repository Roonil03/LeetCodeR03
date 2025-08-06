func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	col := n - 1
	for row := 0; row < m; row++ {
		for col >= 0 && grid[row][col] < 0 {
			col--
		}
		res += n - (col + 1)
	}
	return res
}