func countKthRoots(l int, r int, k int) int {
	if k == 1 {
		return r - l + 1
	}
	return h1(r, k) - h1(l-1, k)
}

func h1(n, k int) int {
	if n >= 0 {
		if n == 0 {
			return 0
		}
		x := int(math.Pow(float64(n), 1.0/float64(k)))
		for a1(x+1, k, n) {
			x++
		}
		for !a1(x, k, n) {
			x--
		}
		return x
	}
	if k%2 == 0 {
		return -1
	}
	a := -n
	b := int(math.Pow(float64(a), 1.0/float64(k)))
	for !a2(b, k, n) {
		b++
	}
	for b > 1 && a2(b-1, k, a) {
		b--
	}
	return -b
}

func a1(x, k, n int) bool {
	p := 1
	for i := 0; i < k; i++ {
		if x > 0 && p > n/x {
			return false
		}
		p *= x
	}
	return p <= n
}

func a2(x, k, n int) bool {
	p := 1
	for i := 0; i < k; i++ {
		if x > 0 && p > n/x {
			return true
		}
		p *= x
	}
	return p >= n
}