func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		midV := matrix[mid/n][mid%n]
		if midV == target {
			return true
		} else if midV < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}