func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	p := make([]int, n)
	for i := range dp {
		dp[i] = 1
		p[i] = -1
	}
	l, id := 1, 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				p[i] = j
				if dp[i] > l {
					l = dp[i]
					id = i
				}
			}
		}
	}
	res := make([]int, 0, l)
	for id != -1 {
		res = append(res, nums[id])
		id = p[id]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}