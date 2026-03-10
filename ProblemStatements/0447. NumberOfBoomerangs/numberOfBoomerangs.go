func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	res := 0
	for i := range points {
		mp := make(map[int]int)
		for j := range points {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			d := dx*dx + dy*dy
			mp[d]++
		}
		for _, v := range mp {
			res += v * (v - 1)
		}
	}
	return res
}