func containsCycle(grid [][]byte) bool {
	n := len(grid)
	vis := make([][]bool, n)
	m := len(grid[0])
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(r, c, pr, pc int, ch byte) bool
	dfs = func(r, c, pr, pc int, ch byte) bool {
		vis[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && grid[nr][nc] == ch {
				if vis[nr][nc] {
					if nr != pr || nc != pc {
						return true
					}
					continue
				}
				if dfs(nr, nc, r, c, ch) {
					return true
				}
			}
		}
		return false
	}
	for i := range n {
		for j := range m {
			if !vis[i][j] {
				if dfs(i, j, -1, -1, grid[i][j]) {
					return true
				}
			}
		}
	}
	return false
}