func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	res := make([]int, n)
	index := 0
	for _, num := range nums {
		if num != 0 {
			res[index] = num
			index++
		}
	}
	return res
}