func isMiddleElementUnique(nums []int) bool {
	m := nums[len(nums)/2]
	for _, v := range nums[:(len(nums) / 2)] {
		if v == m {
			return false
		}
	}
	for _, v := range nums[(len(nums)/2 + 1):] {
		if v == m {
			return false
		}
	}
	return true
}