func minMoves2(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		res += nums[j] - nums[i]
	}
	return res
}