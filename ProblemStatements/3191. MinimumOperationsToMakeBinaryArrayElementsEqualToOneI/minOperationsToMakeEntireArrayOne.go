func minOperations(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			res++
		}
	}
	if nums[n-2] != 1 || nums[n-1] != 1 {
		return -1
	}
	return res
}