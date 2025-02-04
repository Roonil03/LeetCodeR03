func maxAscendingSum(nums []int) int {
	a, b := 0, 0
	a = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			a += nums[i]
		} else {
			b = max(b, a)
			a = nums[i]
		}
	}
	return max(b, a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}