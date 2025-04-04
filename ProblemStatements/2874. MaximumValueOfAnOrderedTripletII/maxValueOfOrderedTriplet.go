func maximumTripletValue(nums []int) int64 {
	a, b := 0, 0
	res := int64(0)
	for _, i := range nums {
		res = max(res, int64(b)*int64(i))
		b = max(b, a-i)
		a = max(a, i)
	}
	return res
}