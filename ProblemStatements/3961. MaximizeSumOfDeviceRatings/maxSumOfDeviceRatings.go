func maxRatings(units [][]int) int64 {
	// if len(units) == 0 || len(units[0]) == 0{
	//     return 0
	// }
	n := len(units[0])
	var s1, s2 int64
	m1, m2 := int64(math.MaxInt64), int64(math.MaxInt64)
	for _, v := range units {
		var a1, a2 int64 = math.MaxInt64, math.MaxInt64
		for _, v1 := range v {
			val := int64(v1)
			if val < a1 {
				a2, a1 = a1, val
			} else if val < a2 {
				a2 = val
			}
		}
		if n == 1 {
			a2 = 0
		}
		s1 += a1
		s2 += a2
		if a2 < m2 {
			m2 = a2
		}
		if a1 < m1 {
			m1 = a1
		}
	}
	res := s2 - m2 + m1
	return max(s1, res)
}