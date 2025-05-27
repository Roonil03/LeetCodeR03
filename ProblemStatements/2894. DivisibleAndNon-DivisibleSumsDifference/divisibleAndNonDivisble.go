func differenceOfSums(n int, m int) int {
	sum := n * (n + 1) / 2
	count := n / m
	div := m * count * (count + 1) / 2
	// num2 := div
	// num1 := sum - div
	return sum - 2*div
}