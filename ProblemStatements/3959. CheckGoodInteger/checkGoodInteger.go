func checkGoodInteger(n int) bool {
	dif := 0
	for n > 0 {
		d := n % 10
		if d >= 8 {
			return true
		}
		dif += d * (d - 1)
		if dif >= 50 {
			return true
		}
		n /= 10
	}
	return false
}