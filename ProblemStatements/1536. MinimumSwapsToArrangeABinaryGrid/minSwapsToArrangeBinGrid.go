func minSwaps(grid [][]int) int {
	n := len(grid)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				pos[i] = j
				break
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		j := -1
		for k := i; k < n; k++ {
			if pos[k] <= i {
				j = k
				break
			}
		}
		if j == -1 {
			return -1
		}
		res += j - i
		for k := j; k > i; k-- {
			pos[k], pos[k-1] = pos[k-1], pos[k]
		}
	}
	return res
}