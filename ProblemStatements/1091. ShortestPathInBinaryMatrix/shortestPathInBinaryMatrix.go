func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	dirs := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	q := list.New()
	q.PushBack([3]int{0, 0, 1})
	vis[0][0] = true
	for q.Len() > 0 {
		e := q.Remove(q.Front()).([3]int)
		r, c, d := e[0], e[1], e[2]
		if r == n-1 && c == n-1 {
			return d
		}
		for _, d1 := range dirs {
			nr, nc := r+d1[0], c+d1[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc] && grid[nr][nc] == 0 {
				vis[nr][nc] = true
				q.PushBack([3]int{nr, nc, d + 1})
			}
		}
	}
	return -1
}