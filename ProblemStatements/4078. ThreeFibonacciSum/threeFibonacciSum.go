func threeFibonacciSum(n int) bool {
	if n < 2 || n&1 != 0 {
		return false
	}
	temp := n >> 1
	a, b := 1, 1
	for b < temp {
		a, b = b, a+b
	}
	return b == temp
}