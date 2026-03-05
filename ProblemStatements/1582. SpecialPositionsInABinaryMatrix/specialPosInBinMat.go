func numSpecial(mat [][]int) int {
	r, c := make([]int, len(mat)), make([]int, len(mat[0]))
	for i := range mat {
		for j := range mat[i] {
			r[i] += mat[i][j]
			c[j] += mat[i][j]
		}
	}
	res := 0
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && r[i] == 1 && c[j] == 1 {
				res++
			}
		}
	}
	return res
}