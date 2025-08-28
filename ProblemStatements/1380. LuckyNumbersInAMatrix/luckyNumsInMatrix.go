func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	a := make([]int, m)
	for i := 0; i < m; i++ {
		temp := matrix[i][0]
		for j := 1; j < n; j++ {
			if matrix[i][j] < temp {
				temp = matrix[i][j]
			}
		}
		a[i] = temp
	}
	b := make([]int, n)
	for j := 0; j < n; j++ {
		temp := matrix[0][j]
		for i := 1; i < m; i++ {
			if matrix[i][j] > temp {
				temp = matrix[i][j]
			}
		}
		b[j] = temp
	}
	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == a[i] && matrix[i][j] == b[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}