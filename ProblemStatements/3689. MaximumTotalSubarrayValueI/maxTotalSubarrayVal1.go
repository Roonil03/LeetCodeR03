func maxTotalValue(nums []int, k int) int64 {
	a, b := slices.Max(nums), slices.Min(nums)
	temp := int64(a - b)
	return temp * int64(k)
}