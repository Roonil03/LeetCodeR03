func maximumTripletValue(nums []int) int64 {
	// var a, b, c int64
	// a, b, c = 0, 0, 0
	// for _, i := range nums{
	//     a = max(a, c * int64(i))
	//     c = max(c, b - int64(i))
	//     b = max(b, int64(i))
	// }
	// return a
	l := len(nums)
	m := int64(0)
	for i := range l - 2 {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				c := int64((nums[i] - nums[j]) * nums[k])
				if c > m {
					m = c
				}
			}
		}
	}
	return m
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}