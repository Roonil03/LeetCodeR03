func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	res := -1
	for _, n := range arr {
		a := dp[n-difference]
		dp[n] = a + 1
		if dp[n] > res {
			res = dp[n]
		}
	}
	return res
}