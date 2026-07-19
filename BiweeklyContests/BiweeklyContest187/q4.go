func minCost(source string, target string, rules [][]string, costs []int) int {
	n := len(source)
	if n != len(target) {
		return -1
	}
	type trans struct {
		l, c int
	}
	adj := make([][]trans, n)
	for i, v := range rules {
		p, r := []byte(v[0]), []byte(v[1])
		l := len(p)
		c := costs[i]
		for j := range l {
			if p[j] == '*' {
				c++
			}
		}
		mat := RabinKarp([]byte(target), r, 257, func(b byte) uint64 {
			return uint64(b)
		})
		for _, id := range mat {
			fg := true
			for j := range l {
				if p[j] != '*' && p[j] != source[id+j] {
					fg = false
					break
				}
			}
			if fg {
				adj[id] = append(adj[id], trans{l, c})
			}
		}
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = /*math.MaxInt*/ (1 << 60)
	}
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		if source[i] == target[i] {
			dp[i] = dp[i+1]
		}
		for _, v := range adj[i] {
			if dp[i+v.l]+v.c < dp[i] {
				dp[i] = dp[i+v.l] + v.c
			}
		}
	}
	if dp[0] >= /*math.MaxInt*/ (1 << 60) {
		return -1
	}
	return dp[0]
}

func RabinKarp[T comparable](text, pattern []T, base uint64, val func(T) uint64) []int {
	n := len(text)
	m := len(pattern)
	if m == 0 || n < m {
		return nil
	}

	var pHash, tHash uint64
	var h uint64 = 1

	for i := 0; i < m-1; i++ {
		h *= base
	}

	for i := 0; i < m; i++ {
		pHash = pHash*base + val(pattern[i])
		tHash = tHash*base + val(text[i])
	}

	var matches []int
	for i := 0; i <= n-m; i++ {
		if pHash == tHash {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				matches = append(matches, i)
			}
		}
		if i < n-m {
			tHash = tHash - val(text[i])*h
			tHash = tHash*base + val(text[i+m])
		}
	}
	return matches
}