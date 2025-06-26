func diagonalSum(mat [][]int) int {
	n := len(mat)
	s := 0
	for i := 0; i < n; i++ {
		s += mat[i][i]
		if i != n-1-i {
			s += mat[i][n-1-i]
		}
	}
	return s
}