func numTilings(n int) int {
	mod := int(1e9 + 7)
	v := make([]int, 1001)
	v[1], v[2], v[3] = 1, 2, 5
	if n <= 3 {
		return v[n]
	}
	for i := 4; i <= n; i++ {
		v[i] = ((2*v[i-1] + v[i-3]) % mod)
	}
	return v[n]
}