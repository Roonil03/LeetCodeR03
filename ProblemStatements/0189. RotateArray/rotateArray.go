func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, s, e int) {
	for s < e {
		temp := nums[s]
		nums[s] = nums[e]
		nums[e] = temp
		s++
		e--
	}
}