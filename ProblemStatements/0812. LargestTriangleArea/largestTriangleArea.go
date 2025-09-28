func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	res := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				area := h1(points[i], points[j], points[k])
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func h1(p1, p2, p3 []int) float64 {
	x1, y1 := float64(p1[0]), float64(p1[1])
	x2, y2 := float64(p2[0]), float64(p2[1])
	x3, y3 := float64(p3[0]), float64(p3[1])
	res := 0.5 * math.Abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))
	return res
}