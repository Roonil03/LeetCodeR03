func interleaveCharacters(word1 string, word2 string, target string) int {
	mod := int64(1e9 + 7)
	w1, w2, t := len(word1), len(word2), len(target)
	dp := make([][]int64, w1+1)
	ndp := make([][]int64, w1+1)
	for i := 0; i <= w1; i++ {
		dp[i] = make([]int64, w2+1)
		ndp[i] = make([]int64, w2+1)
		for j := range dp[i] {
			dp[i][j] = int64(1)
		}
	}
	for i := 1; i <= t; i++ {
		for j := range dp {
			for k := range dp[j] {
				ndp[j][k] = int64(0)
			}
		}
		for j := range dp[0] {
			f1 := int64(0)
			for k := 1; k <= w1; k++ {
				if word1[k-1] == target[i-1] {
					f1 = (f1 + dp[k-1][j]) % mod
				}
				ndp[k][j] = f1
			}
		}
		for j := range dp {
			f2 := int64(0)
			for k := 1; k <= w2; k++ {
				if word2[k-1] == target[i-1] {
					f2 = (f2 + dp[j][k-1]) % mod
				}
				ndp[j][k] = (ndp[j][k] + f2) % mod
			}
		}
		dp, ndp = ndp, dp
	}
	tot := dp[w1][w2]
	a, na := make([]int64, w1+1), make([]int64, w1+1)
	for i := range a {
		a[i] = int64(1)
	}
	for i := 1; i <= t; i++ {
		na[0] = int64(0)
		for j := 1; j <= w1; j++ {
			na[j] = na[j-1]
			if word1[j-1] == target[i-1] {
				na[j] = (na[j] + a[j-1]) % mod
			}
		}
		a, na = na, a
	}
	b, nb := make([]int64, w2+1), make([]int64, w2+1)
	for i := range b {
		b[i] = int64(1)
	}
	for i := 1; i <= t; i++ {
		nb[0] = int64(0)
		for j := 1; j <= w2; j++ {
			nb[j] = nb[j-1]
			if word2[j-1] == target[i-1] {
				nb[j] = (nb[j] + b[j-1]) % mod
			}
		}
		b, nb = nb, b
	}
	res := (tot - a[w1] - b[w2] + 2*mod) % mod
	return int(res)
}