func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		m := l + (r-l)/2
		sq := m * m
		if sq == num {
			return true
		} else if sq < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}