func rotatedDigits(n int) int {
	good := [10]bool{false, false, true, false, false, true, true, false, false, true}
	invalid := [10]bool{false, false, false, true, true, false, false, true, false, false}
	res := 0
	for i := 2; i <= n; i++ {
		x := i
		f1, fg := false, true
		for x > 0 {
			if invalid[x%10] {
				fg = false
				break
			}
			if good[x%10] {
				f1 = true
			}
			x /= 10
		}
		if fg && f1 {
			res++
		}
	}
	return res
}