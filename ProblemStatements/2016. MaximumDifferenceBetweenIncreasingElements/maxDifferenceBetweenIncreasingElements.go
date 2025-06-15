func maximumDifference(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	a, res := nums[0], -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > a {
			d := nums[i] - a
			if d > res {
				res = d
			}
		} else if nums[i] < a {
			a = nums[i]
		}
	}
	return res
}