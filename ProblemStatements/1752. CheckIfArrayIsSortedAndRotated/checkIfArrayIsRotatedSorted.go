func check(nums []int) bool {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			res++
		}
	}
	if nums[len(nums)-1] > nums[0] {
		res++
	}
	return res <= 1
}