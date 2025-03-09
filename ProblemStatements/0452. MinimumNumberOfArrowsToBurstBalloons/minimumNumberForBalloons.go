func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	a := points[0][1]
	for _, point := range points {
		if point[0] > a {
			res++
			a = point[1]
		}
	}
	return res
}