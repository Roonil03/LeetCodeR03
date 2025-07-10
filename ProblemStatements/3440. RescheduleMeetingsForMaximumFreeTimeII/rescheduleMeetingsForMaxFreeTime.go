func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	m := make([][2]int, n)
	for i := 0; i < n; i++ {
		m[i] = [2]int{startTime[i], endTime[i]}
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
	// g := make([]int, n+1)
	// g[0] = m[0][0]
	// for i := 1; i < n; i++ {
	//     g[i] = m[i][0] - m[i-1][1]
	// }
	// g[n] = eventTime - m[n-1][1]
	// res := 0
	// for _, a := range g {
	//     res = max(res, a)
	// }
	// for i := 0; i < n; i++ {
	//     r := g[i] + g[i+1]
	//     res = max(res, r)
	// }
	free := make([]int, n+1)
	free[0] = m[0][0]
	for i := 1; i < n; i++ {
		free[i] = m[i][0] - m[i-1][1]
	}
	free[n] = eventTime - m[n-1][1]
	globalMax := 0
	for i := 0; i < n+1; i++ {
		globalMax = max(globalMax, free[i])
	}
	prefix := make([]int, n+1)
	prefix[0] = free[0]
	for i := 1; i < n+1; i++ {
		prefix[i] = max(prefix[i-1], free[i])
	}
	suffix := make([]int, n+1)
	suffix[n] = free[n]
	for i := n - 1; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], free[i])
	}
	ans := globalMax
	for i := 0; i < n; i++ {
		var merged int
		if i == 0 {
			merged = m[1][0]
		} else if i == n-1 {
			merged = eventTime - m[n-2][1]
		} else {
			merged = m[i+1][0] - m[i-1][1]
		}

		d := m[i][1] - m[i][0]
		var low, high int
		if i == 0 {
			low = 0
			high = 1
		} else if i == n-1 {
			low = n - 1
			high = n
		} else {
			low = i
			high = i + 1
		}

		alt := 0
		if low > 0 {
			alt = max(alt, prefix[low-1])
		}
		if high < n && high+1 < n+1 {
			alt = max(alt, suffix[high+1])
		}

		candidate := merged
		if alt < d {
			candidate = max(0, merged-d)
		}
		ans = max(ans, candidate)
	}
	return ans
}