func maxPrimes(n int, s int) []int {
	if n < 2 || s < 2 {
		return []int{}
	}
	res := make([]int, 0)
	sum := 0
	if sum+2 <= s {
		sum += 2
		res = append(res, 2)
	} else {
		return res
	}
	half := (n + 1) / 2
	bits := make([]uint64, (half+63)/64)
	for i := 1; 2*i*i+2*i < half; i++ {
		if bits[i/64]&(uint64(1)<<(i%64)) == 0 {
			p := 2*i + 1
			for j := 2*i*i + 2*i; j < half; j += p {
				bits[j/64] |= uint64(1) << (j % 64)
			}
		}
	}
	for i := 1; i < half; i++ {
		if bits[i/64]&(uint64(1)<<(i%64)) == 0 {
			p := 2*i + 1
			if sum+p <= s {
				sum += p
				res = append(res, p)
			} else {
				break
			}
		}
	}
	return res
}