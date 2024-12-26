func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum < target || (target+sum)%2 != 0 {
		return 0
	}
	newTarget := (target + sum) / 2
	if newTarget < 0 {
		return 0
	}
	return subsetSum(nums, newTarget)
}

func subsetSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}