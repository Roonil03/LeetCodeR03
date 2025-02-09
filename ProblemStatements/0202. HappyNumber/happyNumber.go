func isHappy(n int) bool {
	a := make(map[int]bool)
	for n != 1 && !a[n] {
		a[n] = true
		n = sum(n)
	}
	return n == 1
}

func sum(n int) int {
	res := 0
	for n > 0 {
		d := n % 10
		res += d * d
		n /= 10
	}
	return res
}