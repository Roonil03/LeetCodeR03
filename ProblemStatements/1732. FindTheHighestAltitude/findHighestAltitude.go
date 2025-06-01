func largestAltitude(gain []int) int {
	res := 0
	a := 0
	for _, g := range gain {
		a += g
		if a > res {
			res = a
		}
	}
	return res
}