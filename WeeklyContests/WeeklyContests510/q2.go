func minimumCost(nums []int, k int) int {
	mod := int64(1e9 + 7)
	m, s := int64(0), int64(0)
	for _, v := range nums {
		s += int64(v)
		if s > int64(k) {
			if a := (s - 1) / int64(k); a > m {
				m = a
			}
		}
	}
	n := m % mod
	return int((n * (n + 1)) % (mod * 2) / 2)
}