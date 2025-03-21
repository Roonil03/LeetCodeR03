func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for a := 0; a <= n-l; a++ {
			b := a + l - 1
			if l == 2 {
				dp[a][b] = s[a] == s[b]
			} else {
				dp[a][b] = s[a] == s[b] && dp[a+1][b-1]
			}
		}
	}
	res := [][]string{}
	c := []string{}
	helper(s, 0, &c, &res, dp)
	return res
}

func helper(s string, a int, c *[]string, res *[][]string, dp [][]bool) {
	if a >= len(s) {
		temp := make([]string, len(*c))
		copy(temp, *c)
		*res = append(*res, temp)
		return
	}
	for b := a; b < len(s); b++ {
		if dp[a][b] {
			*c = append(*c, s[a:b+1])
			helper(s, b+1, c, res, dp)
			*c = (*c)[:len(*c)-1]
		}
	}
}