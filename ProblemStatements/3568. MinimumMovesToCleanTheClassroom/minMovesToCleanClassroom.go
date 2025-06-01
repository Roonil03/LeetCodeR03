func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m-k+1)
	for i := 0; i <= m-k; i++ {
		res[i] = make([]int, n-k+1)
	}
	for i := 0; i <= m-k; i++ {
		f := make(map[int]int)
		for r := i; r < i+k; r++ {
			for c := 0; c < k; c++ {
				f[grid[r][c]]++
			}
		}
		res[i][0] = help(f)
		for j := 1; j <= n-k; j++ {
			for r := i; r < i+k; r++ {
				val := grid[r][j-1]
				f[val]--
				if f[val] == 0 {
					delete(f, val)
				}
			}
			for r := i; r < i+k; r++ {
				val := grid[r][j+k-1]
				f[val]++
			}
			res[i][j] = help(f)
		}
	}
	return res
}

func help(f map[int]int) int {
	if len(f) == 1 {
		return 0
	}
	vals := make([]int, 0, len(f))
	for val := range f {
		vals = append(vals, val)
	}
	sort.Ints(vals)
	d := math.MaxInt32
	for i := 1; i < len(vals); i++ {
		temp := vals[i] - vals[i-1]
		if temp < d {
			d = temp
		}
	}
	return d
}