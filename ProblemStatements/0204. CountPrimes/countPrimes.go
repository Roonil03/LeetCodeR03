func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	np := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if !np[i] {
			res++
			for j := i * i; j < n; j += i {
				np[j] = true
			}
		}
	}
	return res
}