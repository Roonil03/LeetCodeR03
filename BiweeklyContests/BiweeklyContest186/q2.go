func maxValidPairSum(nums []int, k int) int {
	m, res := nums[0], nums[0]+nums[k]
	for i := k; i < len(nums); i++ {
		m = max(m, nums[i-k])
		res = max(res, m+nums[i])
	}
	return res
}