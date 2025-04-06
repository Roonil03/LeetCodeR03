func colorTheArray(n int, queries [][]int) []int {
	balls := make([]int, n)
	res := make([]int, len(queries))
	x := 0
	for i, q := range queries {
		if balls[q[0]] == q[1] {
			res[i] = x
			continue
		}
		if q[0] > 0 && balls[q[0]] > 0 && balls[q[0]-1] == balls[q[0]] {
			x--
		}
		if q[0] < n-1 && balls[q[0]] > 0 && balls[q[0]+1] == balls[q[0]] {
			x--
		}
		if q[0] > 0 && balls[q[0]-1] == q[1] {
			x++
		}
		if q[0] < n-1 && balls[q[0]+1] == q[1] {
			x++
		}
		res[i] = x
		balls[q[0]] = q[1]
	}
	return res
}