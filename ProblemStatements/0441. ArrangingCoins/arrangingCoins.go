func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)/2
		coins := m * (m + 1) / 2
		if coins == n {
			return m
		} else if coins < n {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}