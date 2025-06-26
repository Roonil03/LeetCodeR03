func maximumWealth(accounts [][]int) int {
	m := 0
	for _, a := range accounts {
		s := 0
		for _, v := range a {
			s += v
		}
		if s > m {
			m = s
		}
	}
	return m
}