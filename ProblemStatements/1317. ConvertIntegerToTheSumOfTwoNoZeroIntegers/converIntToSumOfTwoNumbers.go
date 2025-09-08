func getNoZeroIntegers(n int) []int {
	a, rem, u := 0, n, 1
	for rem > 1 {
		dig := rem % 10
		rem /= 10
		if dig < 2 {
			a += u << 1
			rem--
		} else {
			a += u
		}
		u *= 10
	}
	return []int{a, n - a}
}