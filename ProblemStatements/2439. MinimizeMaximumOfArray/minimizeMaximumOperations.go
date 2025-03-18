func minimizeArrayValue(nums []int) int {
	res := 0
	sum := int64(0)
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i])
		c := int((sum + int64(i)) / int64(i+1))
		if c > res {
			res = c
		}
	}
	return res
}