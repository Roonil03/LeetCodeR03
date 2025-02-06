func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	al := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		if al {
			return false
		}
		if i > 0 && nums[i-1] > nums[i+1] {
			nums[i+1] = nums[i]
		}
		al = true
	}
	return true
}