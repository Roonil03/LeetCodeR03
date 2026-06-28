func filterOccupiedIntervals(oi [][]int, fs int, fe int) [][]int {
	slices.SortFunc(oi, func(a, b []int) int {
		return a[0] - b[0]
	})
	var m, res [][]int
	c := []int{oi[0][0], oi[0][1]}
	for _, v := range oi[1:] {
		if v[0] <= c[1]+1 {
			c[1] = max(c[1], v[1])
		} else {
			m = append(m, c)
			c = []int{v[0], v[1]}
		}
	}
	m = append(m, c)
	for _, v := range m {
		if v[1] < fs || v[0] > fe {
			res = append(res, v)
		} else {
			if v[0] < fs {
				res = append(res, []int{v[0], fs - 1})
			}
			if v[1] > fe {
				res = append(res, []int{fe + 1, v[1]})
			}
		}
	}
	return res
}