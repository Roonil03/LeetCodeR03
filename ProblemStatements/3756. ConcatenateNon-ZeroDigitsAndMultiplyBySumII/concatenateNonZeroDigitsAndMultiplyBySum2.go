func sumAndMultiply(s string, queries [][]int) []int {
    res := make([]int, len(queries))
    mod := int64(1e9+7)
    sums := make([]int64, len(s))
    sums[0] = int64(s[0] - '0')
    for i := 1; i < len(s); i++{
        sums[i] = (sums[i-1] + int64(s[i] - '0')) % mod
    }
    p10 := make([]int64, len(s)+1)
	nz := make([]int, len(s)+1)
	val := make([]int64, len(s)+1)
	p10[0] = 1
    for i := range s{
        p10[i+1] = (p10[i] * 10) % mod
		nz[i+1] = nz[i]
		if s[i] != '0' {
			nz[i+1]++
			val[nz[i+1]] = (val[nz[i]]*10 + int64(s[i]-'0')) % mod
		}
    }
    for i, q := range queries {
		l, r := q[0], q[1]
		sum := sums[r]
		if l > 0 {
			sum = (sum - sums[l-1] + mod) % mod
		}
		c1, c2 := nz[l], nz[r+1]
		if c1 != c2 {
			v := (val[c2] - (val[c1]*p10[c2-c1])%mod + mod) % mod
			res[i] = int((v * sum) % mod)
		}
	}
    return res
}