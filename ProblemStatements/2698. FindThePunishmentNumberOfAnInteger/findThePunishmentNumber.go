func punishmentNumber(n int) int {
	var cd func(num, tar int) bool
	cd = func(num, tar int) bool {
		if tar < 0 || num < tar {
			return false
		}
		if num == tar {
			return true
		}
		return cd(num/10, tar-(num%10)) || cd(num/100, tar-(num%100)) || cd(num/1000, tar-(num%1000))
	}
	res := 0
	for i := 1; i <= n; i++ {
		sq := i * i
		if cd(sq, i) {
			res += sq
		}
	}
	return res
}