func minimumOperationsToWriteY(grid [][]int) int {
	ans := math.MaxInt64
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if i != k {
				ans = min(ans, helper(grid, i, k))
			}
		}
	}
	return ans
}

func helper(grid [][]int, y, ny int) int {
	ans, n := 0, len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i <= n/2 && (i == j || i+j == n-1)) || (i > n/2 && j == n/2) {
				if y != grid[i][j] {
					ans++
				}
			} else {
				if ny != grid[i][j] {
					ans++
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}