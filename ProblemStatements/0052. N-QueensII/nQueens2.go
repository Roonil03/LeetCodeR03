package main

func totalNQueens(n int) int {
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	count := 0
	var solve func(row int)
	solve = func(row int) {
		if row == n {
			count++
			return
		}
		for col := 0; col < n; col++ {
			d1 := row - col + n - 1
			d2 := row + col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}
			cols[col], diag1[d1], diag2[d2] = true, true, true
			solve(row + 1)
			cols[col], diag1[d1], diag2[d2] = false, false, false
		}
	}
	solve(0)
	return count
}
