func maximumLength(nums []int) int {
	n := len(nums)
	a, b := 0, 0
	for _, v := range nums {
		if v&1 == 0 {
			a++
		} else {
			b++
		}
	}
	res := max(a, b)
	p := nums[0] & 1
	alt := 1
	for i := 1; i < n; i++ {
		c := nums[i] & 1
		if c != p {
			alt++
			p = c
		}
	}
	if alt > res {
		return alt
	}
	return res
}