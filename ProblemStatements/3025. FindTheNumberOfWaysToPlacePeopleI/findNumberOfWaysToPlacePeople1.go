func numberOfPairs(points [][]int) int {
	n := len(points)
	ys := make([]int, n)
	for i, p := range points {
		ys[i] = p[1]
	}
	sort.Ints(ys)
	yc := make(map[int]int)
	id := 1
	for _, y := range ys {
		if _, ex := yc[y]; !ex {
			yc[y] = id
			id++
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	res := 0
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		for j := 0; j < i; j++ {
			px, py := points[j][0], points[j][1]
			if px <= x && py >= y && (px < x || py > y) {
				b := true
				for k := 0; k < n; k++ {
					if k == i || k == j {
						continue
					}
					kx, ky := points[k][0], points[k][1]
					if px <= kx && kx <= x && y <= ky && ky <= py {
						b = false
						break
					}
				}
				if b {
					res++
				}
			}
		}
	}
	return res
}
