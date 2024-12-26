func findTargetSumWays(nums []int, target int) int {
	memo := make(map[[2]int]int)
	var backtrack func(id, cur int) int
	backtrack = func(id, cur int) int {
		if id == len(nums) {
			if cur == target {
				return 1
			}
			return 0
		}
		key := [2]int{id, cur}
		if val, exists := memo[key]; exists {
			return val
		}
		pos := backtrack(id+1, cur+nums[id])
		neg := backtrack(id+1, cur-nums[id])
		memo[key] = pos + neg
		return memo[key]
	}
	return backtrack(0, 0)
}