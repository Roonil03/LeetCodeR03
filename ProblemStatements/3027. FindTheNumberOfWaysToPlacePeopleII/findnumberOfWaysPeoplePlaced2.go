func numberOfPairs(points [][]int) int {
	//     n := len(points)
	//     res := 0
	//     for i := 0; i < n; i++{
	//         for j := 0; j < n; j++{
	//             if i == j{
	//                 continue
	//             }
	//             x1, y1 := points[i][0], points[i][1]
	//             x2, y2 := points[j][0], points[j][1]
	//             if x1 <= x2 && y1 >= y2{
	//                 valid := true
	//                 for k := 0; k < n; k++{
	//                     if k == i || k == j{
	//                         continue
	//                     }
	//                     x3, y3 := points[k][0], points[k][1]
	//                     if x1 <= x3 && x3 <= x2 && y2 <= y3 && y3 <= y1{
	//                         valid = false
	//                         break
	//                     }
	//                 }
	//                 if valid{
	//                     res++
	//                 }
	//             }
	//         }
	//     }
	//     return res

	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	cnt := 0
	for i := 0; i < n; i++ {
		top := points[i][1]
		bot := -1 << 31
		for j := i + 1; j < n; j++ {
			y := points[j][1]
			if bot < y && y <= top {
				cnt++
				bot = y
				if bot == top {
					break
				}
			}
		}
	}
	return cnt

}