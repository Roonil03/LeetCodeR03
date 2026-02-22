func isDigitorialPermutation(n int) bool {
	fact := make([]int, 10)
	fact[0] = 1
	for i := 1; i < 10; i++ {
		fact[i] = fact[i-1] * i
	}
	count := [10]int{}
	sum := 0
	temp := n
	for temp > 0 {
		dig := temp % 10
		count[dig]++
		sum += fact[dig]
		temp /= 10
	}
	t := [10]int{}
	s := sum
	for s > 0 {
		dig := s % 10
		t[dig]++
		s /= 10
	}
	return count == t
}