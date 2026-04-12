func findDegrees(matrix [][]int) []int {
	n := len(matrix)
	res := make([]int, n)
	for i := range n {
		deg := 0
		for j := range n {
			deg += matrix[i][j]
		}
		res[i] = deg
	}
	return res
}