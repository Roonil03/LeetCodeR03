func sumOfGoodNumbers(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		isG := true
		if i-k >= 0 {
			if nums[i] <= nums[i-k] {
				isG = false
			}
		}
		if i+k < len(nums) {
			if nums[i] <= nums[i+k] {
				isG = false
			}
		}
		if isG {
			res += nums[i]
		}
	}
	return res
}