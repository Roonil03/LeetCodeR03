func longestMonotonicSubarray(nums []int) int {
	a, b := 0, 0
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			a++
			b = 0
		} else if nums[i-1] < nums[i] {
			b++
			a = 0
		} else {
			a = 0
			b = 0
		}
		res = max(max(a, b), res)
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}