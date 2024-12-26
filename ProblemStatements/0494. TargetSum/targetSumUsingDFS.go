func findTargetSumWays(nums []int, target int) int {
	var dfs func(id, cur int) int
	dfs = func(id, cur int) int {
		if id == len(nums) {
			if cur == target {
				return 1
			}
			return 0
		}
		return dfs(id+1, cur+nums[id]) + dfs(id+1, cur-nums[id])
	}
	return dfs(0, 0)
}