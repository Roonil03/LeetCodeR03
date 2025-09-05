func kClosest(points [][]int, k int) [][]int {
	type h struct {
		x, y, dist int
	}
	arr := make([]h, len(points))
	for i, p := range points {
		arr[i] = h{
			x:    p[0],
			y:    p[1],
			dist: p[0]*p[0] + p[1]*p[1],
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].dist < arr[j].dist
	})
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = []int{arr[i].x, arr[i].y}
	}
	return res
}