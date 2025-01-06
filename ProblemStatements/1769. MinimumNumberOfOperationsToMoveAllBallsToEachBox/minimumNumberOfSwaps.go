func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	l, r := 0, 0
	lo, ro := 0, 0
	for i := 0; i < n; i++ {
		if boxes[i] == '1' {
			r++
			ro += i
		}
	}
	for i := 0; i < n; i++ {
		res[i] = lo + ro
		if boxes[i] == '1' {
			l++
			r--
		}
		lo += l
		ro -= r
	}
	return res
}