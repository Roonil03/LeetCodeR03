func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    dist:= make([][]int, n)
    var q [][2]int
    for i := range n{
        dist[i] = make([]int, n)
        for j := range n{
            dist[i][j] = -1
            if grid[i][j] == 1{
                q = append(q, [2]int{i, j})
                dist[i][j] = 0
            }
        }
    }
    dir := []int{-1, 0, 1, 0, -1}
    for len(q) > 0{
        p := q[0]
        q = q[1:]
        for i := range 4{
            r, c := p[0] + dir[i], p[1] + dir[i + 1]
            if r >= 0&& r< n && c >= 0 && c < n && dist[r][c] == -1{
                dist[r][c] = dist[p[0]][p[1]] + 1
                q = append(q, [2]int{r, c})
            }
        }
    }
    l, r := 0, n * 2
    res := 0
    for l <= r{
        mid := (l + r) / 2
        if dist[0][0] < mid || dist[n - 1][n - 1] < mid{
            r = mid - 1
            continue
        }
        vis := make([][]bool, n)
        for i := range n{
            vis[i] = make([]bool, n)
        }
        vq := [][2]int{{0, 0}}
        fg := false
        for len(vq) > 0{
            p := vq[0]
            vq = vq[1:]
            if p[0] == n - 1 && p[1] == n - 1{
                fg = true
                break
            }
            for i := range 4{
                nr, nc := p[0] + dir[i], p[1] + dir[i + 1]
                if nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc] && dist[nr][nc] >= mid{
                    vis[nr][nc] = true
                    vq = append(vq, [2]int{nr, nc})
                }
            }
        }
        if fg{
            res = mid
            l = mid + 1
        } else{
            r = mid - 1
        }
    }
    return res
}