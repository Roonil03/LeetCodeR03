func almostPalindromic(s string) int {
	res := 0
	n := len(s)
	if isPal(s, n) {
		return n
	}
	pal := make([][]bool, n)
	dp := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
		dp[i] = make([]bool, n)
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if l == 1 {
				pal[i][j] = true
			} else if l == 2 {
				pal[i][j] = (s[i] == s[j])
			} else {
				pal[i][j] = (s[i] == s[j] && pal[i+1][j-1])
			}
			if l >= 2 {
				if pal[i+1][j] || pal[i][j-1] {
					dp[i][j] = true
				} else if s[i] == s[j] && dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] {
				res = max(res, l)
			}
		}
	}
	return res
}

func isPal(s string, n int) bool {
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		if s[i] != s[j] {
			return false
		}
	}
	return true
}