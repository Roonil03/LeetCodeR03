func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	states := [][][]int{
		{},
		{{0, -1}, {0, 1}},
		{{-1, 0}, {1, 0}},
		{{0, -1}, {1, 0}},
		{{0, 1}, {1, 0}},
		{{0, -1}, {-1, 0}},
		{{0, 1}, {-1, 0}},
	}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	q := [][2]int{{0, 0}}
	vis[0][0] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		r, c := cur[0], cur[1]
		if r == m-1 && c == n-1 {
			return true
		}
		for _, d := range states[grid[r][c]] {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !vis[nr][nc] {
				a, b := -d[0], -d[1]
				if h1(grid[nr][nc], a, b, states) {
					vis[nr][nc] = true
					q = append(q, [2]int{nr, nc})
				}
			}
		}
	}
	return false
}

func h1(t, r, c int, states [][][]int) bool {
	for _, v := range states[t] {
		if v[0] == r && v[1] == c {
			return true
		}
	}
	return false
}