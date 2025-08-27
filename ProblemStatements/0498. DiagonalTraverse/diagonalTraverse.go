func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	res := []int{}
	dirs := [2][2]int{{-1, 1}, {1, -1}}
	r, c, d := 0, 0, 0
	for i := 0; i < m*n; i++ {
		res = append(res, mat[r][c])
		r += dirs[d][0]
		c += dirs[d][1]
		if r >= m {
			r = m - 1
			c += 2
			d ^= 1
		} else if c >= n {
			c = n - 1
			r += 2
			d ^= 1
		} else if r < 0 {
			r = 0
			d ^= 1
		} else if c < 0 {
			c = 0
			d ^= 1
		}
	}
	return res
}