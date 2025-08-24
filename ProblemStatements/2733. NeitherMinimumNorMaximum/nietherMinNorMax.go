func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	a, b := nums[0], nums[0]
	for _, i := range nums {
		if i < a {
			a = i
		}
		if i > b {
			b = i
		}
	}
	for _, i := range nums {
		if i != a && i != b {
			return i
		}
	}
	return -1
}