func gcdSum(nums []int) int64 {
	m := nums[0]
	p := make([]int, len(nums))
	for i, v := range nums {
		if v > m {
			m = v
		}
		a, b := v, m
		for b != 0 {
			a, b = b, a%b
		}
		p[i] = a
	}
	sort.Ints(p)
	res := int64(0)
	for i := range len(p) / 2 {
		a, b := p[i], p[len(p)-1-i]
		for b != 0 {
			a, b = b, a%b
		}
		res += int64(a)
	}
	return res
}