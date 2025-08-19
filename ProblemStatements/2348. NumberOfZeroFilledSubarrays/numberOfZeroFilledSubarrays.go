func zeroFilledSubarray(nums []int) int64 {
	a, res := int64(0), int64(0)
	for _, i := range nums {
		if i == 0 {
			a++
		} else {
			a = 0
		}
		res += a
	}
	return res
}