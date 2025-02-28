func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	var res []byte
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			res = append(res, str1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			res = append(res, str1[i-1])
			i--
		} else {
			res = append(res, str2[j-1])
			j--
		}
	}
	for i > 0 {
		res = append(res, str1[i-1])
		i--
	}
	for j > 0 {
		res = append(res, str2[j-1])
		j--
	}
	reverse(res)
	return string(res)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}