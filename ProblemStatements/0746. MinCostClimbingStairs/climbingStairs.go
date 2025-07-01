func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	dp1, dp2 := cost[1], cost[0]
	for i := 2; i < n; i++ {
		curr := cost[i] + min(dp1, dp2)
		dp2 = dp1
		dp1 = curr
	}
	return min(dp1, dp2)
}