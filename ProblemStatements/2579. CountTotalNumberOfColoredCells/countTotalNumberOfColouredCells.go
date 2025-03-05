func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}
	return int64(1 + 4*(n-1)*n/2)
}