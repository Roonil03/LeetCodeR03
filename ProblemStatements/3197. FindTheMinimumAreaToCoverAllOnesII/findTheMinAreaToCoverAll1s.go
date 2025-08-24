func minimumSum(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	res := math.MaxInt32
	h1 := func(r1, r2, c1, c2 int) int {
		a1, a2 := math.MaxInt32, -1
		b1, b2 := math.MaxInt32, -1
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				if grid[i][j] == 1 {
					if i < a1 {
						a1 = i
					}
					if i > a2 {
						a2 = i
					}
					if j < b1 {
						b1 = j
					}
					if j > b2 {
						b2 = j
					}
				}
			}
		}
		if a1 == math.MaxInt32 {
			return 0
		}
		return (a2 - a1 + 1) * (b2 - b1 + 1)
	}
	for i := 0; i < r-2; i++ {
		for j := i + 1; j < r-1; j++ {
			area1 := h1(0, i, 0, c-1)
			area2 := h1(i+1, j, 0, c-1)
			area3 := h1(j+1, r-1, 0, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	for i := 0; i < c-2; i++ {
		for j := i + 1; j < c-1; j++ {
			area1 := h1(0, r-1, 0, i)
			area2 := h1(0, r-1, i+1, j)
			area3 := h1(0, r-1, j+1, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	for i := 0; i < r-1; i++ {
		for j := 0; j < c-1; j++ {
			area1 := h1(0, i, 0, j)
			area2 := h1(0, i, j+1, c-1)
			area3 := h1(i+1, r-1, 0, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 0; j < c-1; j++ {
			area1 := h1(0, i-1, 0, c-1)
			area2 := h1(i, r-1, 0, j)
			area3 := h1(i, r-1, j+1, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	for i := 0; i < c-1; i++ {
		for j := 0; j < r-1; j++ {
			area1 := h1(0, j, 0, i)
			area2 := h1(j+1, r-1, 0, i)
			area3 := h1(0, r-1, i+1, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	for i := 1; i < c; i++ {
		for j := 0; j < r-1; j++ {
			area1 := h1(0, r-1, 0, i-1)
			area2 := h1(0, j, i, c-1)
			area3 := h1(j+1, r-1, i, c-1)
			if area1 > 0 && area2 > 0 && area3 > 0 {
				total := area1 + area2 + area3
				if total < res {
					res = total
				}
			}
		}
	}
	return res
}