func shortestBridge(grid [][]int) int {
	n := len(grid)
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	q := [][]int{}
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, n)
	}
	ok := false
	for i := 0; i < n && !ok; i++ {
		for j := 0; j < n && !ok; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, &q, vis, dir)
				ok = true
			}
		}
	}
	lvl := 0
	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			cur := q[0]
			q = q[1:]
			r, c := cur[0], cur[1]
			for _, d := range dir {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc] {
					if grid[nr][nc] == 1 {
						return lvl
					}
					vis[nr][nc] = true
					q = append(q, []int{nr, nc})
				}
			}
		}
		lvl++
	}
	return -1
}

func dfs(grid [][]int, r, c int, q *[][]int, vis [][]bool, dir [][]int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || vis[r][c] || grid[r][c] == 0 {
		return
	}
	vis[r][c] = true
	*q = append(*q, []int{r, c})
	for _, d := range dir {
		dfs(grid, r+d[0], c+d[1], q, vis, dir)
	}
}