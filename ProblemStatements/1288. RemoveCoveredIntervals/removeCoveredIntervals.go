func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res, temp := 0, 0
	for _, v := range intervals {
		if v[1] > temp {
			res++
			temp = v[1]
		}
	}
	return res
}