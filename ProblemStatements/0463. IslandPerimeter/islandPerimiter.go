func islandPerimeter(grid [][]int) int {
	res := 0
	for i, v := range grid {
		for j, w := range v {
			if w == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j > 0 && v[j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}