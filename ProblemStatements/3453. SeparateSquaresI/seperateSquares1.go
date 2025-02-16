func separateSquares(squares [][]int) float64 {
	n := len(squares)
	t := 0.0
	minY := math.MaxFloat64
	maxY := -math.MaxFloat64
	for i := 0; i < n; i++ {
		y := float64(squares[i][1])
		l := float64(squares[i][2])
		t += l * l
		if y < minY {
			minY = y
		}
		if y+l > maxY {
			maxY = y + l
		}
	}
	tar := t / 2
	lo, hi := minY, maxY
	for i := 0; i < 100; i++ {
		mid := (lo + hi) / 2
		sum := 0.0
		for j := 0; j < n; j++ {
			y := float64(squares[j][1])
			l := float64(squares[j][2])
			top := y + l
			if mid <= y {
				sum += l * l
			} else if mid < top {
				sum += (top - mid) * l
			}
		}
		if sum > tar {
			lo = mid
		} else {
			hi = mid
		}
	}
	return hi
}