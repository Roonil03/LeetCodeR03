func minimumArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	a1, a2 := n, -1
	b1, b2 := m, -1
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 {
				if r < a1 {
					a1 = r
				}
				if r > a2 {
					a2 = r
				}
				if c < b1 {
					b1 = c
				}
				if c > b2 {
					b2 = c
				}
			}
		}
	}
	if a1 == n {
		return 0
	}
	return (a2 - a1 + 1) * (b2 - b1 + 1)
}