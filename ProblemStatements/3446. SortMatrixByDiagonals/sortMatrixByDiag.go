func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	h1 := func(r, c int, asc bool) {
		a, b := r, c
		vals := []int{}
		for a < n && b < n {
			vals = append(vals, grid[a][b])
			a++
			b++
		}
		if asc {
			sort.Ints(vals)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(vals)))
		}
		a, b, i := r, c, 0
		for a < n && b < n {
			grid[a][b] = vals[i]
			a++
			b++
			i++
		}
	}
	for r := 0; r < n; r++ {
		h1(r, 0, false)
	}
	for c := 1; c < n; c++ {
		h1(0, c, true)
	}
	return grid
}