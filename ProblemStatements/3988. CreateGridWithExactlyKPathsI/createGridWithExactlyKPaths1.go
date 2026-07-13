func createGrid(m int, n int, k int) []string {
	g := make([][]byte, m)
	for i := range g {
		g[i] = make([]byte, n)
		for j := range g[i] {
			g[i][j] = '.'
		}
	}
	dp := make([]int, n)
	h1 := func() int {
		clear(dp)
		dp[0] = 1
		for i := range m {
			for j := range n {
				if g[i][j] == '#' {
					dp[j] = 0
				} else if j > 0 {
					dp[j] += dp[j-1]
				}
			}
		}
		return dp[n-1]
	}
	c := h1()
	if c < k {
		return []string{}
	}
	if c == k {
		return h2(g)
	}
	for i := range m {
		for j := range n {
			if (i == 0 && j == 0) || (i == m-1 && j == n-1) {
				continue
			}
			g[i][j] = '#'
			c = h1()
			if c == k {
				return h2(g)
			}
			if c < k {
				g[i][j] = '.'
			}
		}
	}
	return []string{}
}

func h2(a [][]byte) []string {
	res := make([]string, len(a))
	for i := range a {
		res[i] = string(a[i])
	}
	return res
}