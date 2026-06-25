func zigZagArrays(n int, l int, r int) int {
	k := r - l + 1
	mod := 1000000007
	// inc, dec := make([]int, k), make([]int, k)
	// for i := range inc{
	//     inc[i], dec[i] = 1, 1
	// }
	// ninc, ndec := make([]int, k), make([]int, k)
	// for i := 2; i <= n; i++{
	//     sdec, sinc := 0, 0
	//     for j := range k{
	//         ninc[j] = sdec
	// 		sdec += dec[j]
	// 		if sdec >= mod {
	// 			sdec -= mod
	// 		}
	// 	}
	// 	for j := k - 1; j >= 0; j-- {
	// 		ndec[j] = sinc
	// 		sinc += inc[j]
	// 		if sinc >= mod {
	// 			sinc -= mod
	// 		}
	// 	}
	// 	inc, ninc = ninc, inc
	// 	dec, ndec = ndec, dec
	// }
	// res := 0
	// for i := range k{
	//     res = (res + inc[i] + dec[i]) % mod
	// }
	// return res
	m := make([][]int, 2*k)
	for i := range m {
		m[i] = make([]int, 2*k)
	}
	for i := range k {
		for j := range i {
			m[i][k+j] = 1
		}
		for j := i + 1; j < k; j++ {
			m[k+i][j] = 1
		}
	}
	m = pow1(m, n-1, 2*k, mod)
	res := 0
	for i := range 2 * k {
		for j := range 2 * k {
			res = (res + m[i][j]) % mod
		}
	}
	return res
}

func pow1(a [][]int, p int, s int, mod int) [][]int {
	res := make([][]int, s)
	for i := range res {
		res[i] = make([]int, s)
		res[i][i] = 1
	}
	base := a
	for p > 0 {
		if p&1 == 1 {
			res = h1(res, base, s, mod)
		}
		base = h1(base, base, s, mod)
		p >>= 1
	}
	return res
}

func h1(a, b [][]int, s int, mod int) [][]int {
	c := make([][]int, s)
	for i := range c {
		c[i] = make([]int, s)
	}
	for i := range s {
		for j := range s {
			if a[i][j] == 0 {
				continue
			}
			for k := range s {
				c[i][k] = (c[i][k] + a[i][j]*b[j][k]) % mod
			}
		}
	}
	return c
}