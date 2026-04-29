func maximumScore(grid [][]int) int64 {
	n := len(grid)
	pref := make([][]int64, n)
	for i := range n {
		pref[i] = make([]int64, n+1)
		for j := range n {
			pref[i][j+1] = pref[i][j] + int64(grid[j][i])
		}
	}
	dp := make([][]int64, n+1)
	ndp := make([][]int64, n+1)
	for i := range n + 1 {
		dp[i] = make([]int64, n+1)
		ndp[i] = make([]int64, n+1)
		for j := range n + 1 {
			dp[i][j] = -1
		}
	}
	for j := range n + 1 {
		dp[j][0] = 0
	}
	pMax := make([]int64, n+1)
	sMax := make([]int64, n+1)
	for i := 1; i < n; i++ {
		for r := 0; r <= n; r++ {
			for c := 0; c <= n; c++ {
				ndp[r][c] = -1
			}
		}
		for hPrev := 0; hPrev <= n; hPrev++ {
			maxG := int64(-1)
			for k := 0; k <= n; k++ {
				if dp[hPrev][k] > maxG {
					maxG = dp[hPrev][k]
				}
				pMax[k] = maxG
			}
			maxGPref := int64(-1)
			for k := n; k >= 0; k-- {
				val := int64(-1)
				if dp[hPrev][k] != -1 {
					var gain int64
					if k > hPrev {
						gain = pref[i-1][k] - pref[i-1][hPrev]
					}
					val = dp[hPrev][k] + gain
				}
				if val > maxGPref {
					maxGPref = val
				}
				sMax[k] = maxGPref
			}
			for hCurr := 0; hCurr <= n; hCurr++ {
				var g1 int64
				if hCurr > hPrev {
					g1 = pref[i-1][hCurr] - pref[i-1][hPrev]
				}
				v1 := int64(-1)
				if pMax[hCurr] != -1 {
					v1 = pMax[hCurr] + g1
				}
				v2 := int64(-1)
				if hCurr < n {
					v2 = sMax[hCurr+1]
				}

				res := v1
				if v2 > res {
					res = v2
				}
				ndp[hCurr][hPrev] = res
			}
		}
		dp, ndp = ndp, dp
	}
	var res int64
	for hCurr := 0; hCurr <= n; hCurr++ {
		for hPrev := 0; hPrev <= n; hPrev++ {
			if dp[hCurr][hPrev] == -1 {
				continue
			}
			var gain int64
			if hPrev > hCurr {
				gain = pref[n-1][hPrev] - pref[n-1][hCurr]
			}
			if dp[hCurr][hPrev]+gain > res {
				res = dp[hCurr][hPrev] + gain
			}
		}
	}
	return res
}