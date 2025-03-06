func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	numMap := make(map[int]int)
	var a, b int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			numMap[grid[i][j]]++
		}
	}
	for i := 1; i <= n*n; i++ {
		if numMap[i] == 2 {
			a = i
		} else if numMap[i] == 0 {
			b = i
		}
	}
	return []int{a, b}
}