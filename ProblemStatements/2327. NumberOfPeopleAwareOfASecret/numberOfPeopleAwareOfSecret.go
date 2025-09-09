func peopleAwareOfSecret(n int, delay int, forget int) int {
	mod := int(1e9 + 7)
	dp := make([]int, n+2)
	dp[1] = 1
	s := 0
	for i := 2; i <= n; i++ {
		if i-delay >= 1 {
			s += dp[i-delay]
			if s >= mod {
				s -= mod
			}
		}
		if i-forget >= 1 {
			s -= dp[i-forget]
			if s < 0 {
				s += mod
			}
		}
		dp[i] = s
	}
	res := 0
	st := n - forget + 1
	if st < 1 {
		st = 1
	}
	for i := st; i <= n; i++ {
		res += dp[i]
		if res >= mod {
			res -= mod
		}
	}
	return res
}