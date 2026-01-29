func minCost(grid [][]int, k int) int {
	type cell struct{ r, c, v int }
	cells := make([]cell, 0, len(grid)*len(grid[0]))
	for r := range grid {
		for c := range grid[r] {
			cells = append(cells, cell{r, c, grid[r][c]})
		}
	}
	slices.SortFunc(cells, func(a, b cell) int {
		return cmp.Compare(b.v, a.v)
	})
	dp := make([]int, len(grid)*len(grid[0]))
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	h1 := func(t1 []int) {
		for r := range grid {
			for c := range grid[r] {
				id := r*len(grid[r]) + c
				// if r > 0 && t1[(r - 1) * len(grid[r]) + c] != math.MaxInt{
				//     t1[id] = min(t1[id], t1[(r - 1) * len(grid[r]) + c]) + grid[r][c]
				// }
				// if c > 0 && t1[(r) * len(grid[r]) + (c-1)] != math.MaxInt{
				//     t1[id] = min(t1[id], t1[r * len(grid[r]) + (c - 1)]) + grid[r][c]
				// }
				prev := math.MaxInt
				if r > 0 {
					prev = min(prev, t1[(r-1)*len(grid[0])+c])
				}
				if c > 0 {
					prev = min(prev, t1[r*len(grid[0])+(c-1)])
				}
				if prev != math.MaxInt {
					t1[id] = min(t1[id], prev+grid[r][c])
				}
				// fmt.Print(t1[id], " ")
			}
		}
		// fmt.Println()
	}
	h1(dp)
	res := dp[len(grid)*len(grid[0])-1]
	for t := 1; t <= k; t++ {
		n_dp := make([]int, len(grid)*len(grid[0]))
		for i := range n_dp {
			n_dp[i] = math.MaxInt
		}
		temp := math.MaxInt
		for i := 0; i < len(cells); {
			val := cells[i].v
			id := i
			for i < len(cells) && cells[i].v == val {
				temp = min(temp, dp[cells[i].r*len(grid[0])+cells[i].c])
				i++
			}
			for j := id; j < i; j++ {
				n_dp[cells[j].r*len(grid[0])+cells[j].c] = temp
			}
		}
		h1(n_dp)
		dp = n_dp
		res = min(res, dp[len(grid)*len(grid[0])-1])
	}
	return res
}