func countLargestGroup(n int) int {
	m := make([]int, 37)
	maxCount := 0
	res := 0
	for i := 1; i <= n; i++ {
		num := i
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		m[sum]++
		if maxCount < m[sum] {
			maxCount = m[sum]
			res = 1
		} else if maxCount == m[sum] {
			res++
		}
	}
	return res
}