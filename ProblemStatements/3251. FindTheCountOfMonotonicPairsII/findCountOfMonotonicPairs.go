const mod = 1000000007

func countOfPairs(nums []int) int {
	n := len(nums)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, nums[i]+1)
	}
	for x := 0; x <= nums[0]; x++ {
		dp[0][x] = 1
	}
	for i := 1; i < n; i++ {
		d := 0
		if nums[i] > nums[i-1] {
			d = nums[i] - nums[i-1]
		}
		p := make([]int64, len(dp[i-1]))
		p[0] = dp[i-1][0] % mod
		for j := 1; j < len(dp[i-1]); j++ {
			p[j] = (p[j-1] + dp[i-1][j]) % mod
		}
		for x := 0; x <= nums[i]; x++ {
			if x >= d {
				dp[i][x] = p[x-d]
			} else {
				dp[i][x] = 0
			}
		}
	}
	res := 0
	for _, val := range dp[n-1] {
		res = (res + int(val)) % mod
	}
	return res
}