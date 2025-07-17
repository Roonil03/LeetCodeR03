func maximumLength(nums []int, k int) int {
	//n := len(nums)
	res := 0
	for r := 0; r < k; r++ {
		dp := make([]int, k)
		for _, n := range nums {
			rem := n % k
			p := (r - rem + k) % k
			dp[rem] = max(dp[rem], dp[p]+1)
		}
		for _, l := range dp {
			res = max(res, l)
		}
	}
	return res
}