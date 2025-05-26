func moveZeroes(nums []int) {
	n := len(nums)
	a := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[a] = nums[i]
			a++
		}
	}
	for i := a; i < n; i++ {
		nums[i] = 0
	}
}