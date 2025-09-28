func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}