func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1
	if m <= 0 {
		return 0
	}
	if n == 1 {
		return m
	}
	mod := 1000000007
	inc, dec := make([]int, m), make([]int, m)
	for i := range m {
		inc[i] = 1
		dec[i] = 1
	}
	for i := 1; i < n; i++ {
		ninc := make([]int, m)
		ndec := make([]int, m)
		sdec := 0
		for j := range m {
			ninc[j] = sdec
			sdec = (sdec + dec[j]) % mod
		}
		sinc := 0
		for j := m - 1; j >= 0; j-- {
			ndec[j] = sinc
			sinc = (sinc + inc[j]) % mod
		}
		inc = ninc
		dec = ndec
	}
	res := 0
	for i := range m {
		res = (res + inc[i]) % mod
		res = (res + dec[i]) % mod
	}
	return res
}