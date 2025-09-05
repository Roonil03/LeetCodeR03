func makeTheIntegerZero(num1 int, num2 int) int {
	for k := 1; k <= 60; k++ {
		tar := num1 - k*num2
		if tar >= 0 && pop(tar) <= k && tar >= k {
			return k
		}
	}
	return -1
}

func pop(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}