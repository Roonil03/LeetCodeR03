func productQueries(n int, queries [][]int) []int {
	mod := int(1e9 + 7)
	p := []int{}
	i := 0
	for n > 0 {
		if n&1 == 1 {
			p = append(p, h1(2, i, mod))
		}
		n >>= 1
		i++
	}
	pre := make([]int, len(p)+1)
	pre[0] = 1
	for i := 0; i < len(p); i++ {
		pre[i+1] = (pre[i] * p[i]) % mod
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		num, den := pre[r+1], pre[l]
		inv := h1(den, mod-2, mod)
		res[i] = (num * inv) % mod
	}
	return res
}

func h1(base, exp, mod int) int {
	res := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}