func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	if k == 1 {
		return 0
	}
	temp := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		temp[i] = weights[i] + weights[i+1]
	}
	sort.Ints(temp)
	var m1, m2 int64
	m1 = 0
	for i := 0; i < k-1; i++ {
		m1 += int64(temp[i])
	}
	for i := 0; i < k-1; i++ {
		m2 += int64(temp[n-2-i])
	}
	return m2 - m1
}