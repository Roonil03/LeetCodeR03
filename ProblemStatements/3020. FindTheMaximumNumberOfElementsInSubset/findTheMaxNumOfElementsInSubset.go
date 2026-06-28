func maximumLength(nums []int) int {
	c := make(map[int]int)
	for _, v := range nums {
		c[v]++
	}
	r := c[1]
	if r%2 == 0 {
		r--
	}
	if r < 1 {
		r = 1
	}
	for i := range c {
		if i == 1 {
			continue
		}
		l, v := 0, i
		for c[v] > 1 {
			l += 2
			v *= v
		}
		if c[v] > 0 {
			l++
		} else {
			l--
		}
		if l > r {
			r = l
		}
	}
	return r
}