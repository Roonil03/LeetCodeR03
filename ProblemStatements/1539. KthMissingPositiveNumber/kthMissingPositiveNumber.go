func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		miss := arr[m] - (m + 1)
		if miss < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return k + l
}
