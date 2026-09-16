func numberOfSets(n int, k int) int {
	mod := int(1e9 + 7)
	num, den := 1, 1
	for i := 1; i <= 2*k; i++ {
		num = (num * (n + k - i)) % mod
		den = (den * i) % mod
	}
	p, b := 1, mod-2
	for b > 0 {
		if b&1 == 1 {
			p = (p * den) % mod
		}
		den = (den * den) % mod
		b >>= 1
	}
	return (num * p) % mod
}