func maxSum(nums []int, k int, mul int) int64 {
	n := len(nums)
	s := make([]int, n)
	copy(s, nums)
	slices.Sort(s)
	g1 := func(x, m int) int64 {
		v1 := int64(x)
		v2 := int64(x) * int64(m)
		if v1 > v2 {
			return v1
		}
		return v2
	}
	var cur int64
	for i := range k {
		cur += g1(s[n-k+i], mul-k+1+i)
	}
	res := cur
	for i := range k {
		m := mul - k + 1 + i
		cur += g1(s[i], m) - g1(s[n-k+i], m)
		if cur > res {
			res = cur
		}
	}
	return res
}