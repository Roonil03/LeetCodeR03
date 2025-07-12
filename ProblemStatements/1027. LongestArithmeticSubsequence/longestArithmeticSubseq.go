func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			if v, ok := dp[j][d]; ok {
				dp[i][d] = v + 1
			} else {
				dp[i][d] = 2
			}
			if dp[i][d] > res {
				res = dp[i][d]
			}
		}
	}
	return res
}