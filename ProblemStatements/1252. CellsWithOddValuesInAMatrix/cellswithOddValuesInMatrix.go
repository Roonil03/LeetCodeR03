func oddCells(m int, n int, indices [][]int) int {
	a, b := make([]bool, m), make([]bool, n)
	for _, id := range indices {
		a[id[0]] = !a[id[0]]
		b[id[1]] = !b[id[1]]
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i] != b[j] {
				res++
			}
		}
	}
	return res
}