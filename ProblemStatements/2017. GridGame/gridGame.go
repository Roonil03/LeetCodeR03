// func gridGame(grid [][]int) int64 {
//     n := len(grid[0])
//     top := make([]int64, n)
//     bot := make([]int64, n)
//     top[0] = int64(grid[0][0])
//     bot[0] = int64(grid[1][0])
//     for i := 1; i < n; i++{
//         top[i] = top[i-1] + int64(grid[0][i])
//         bot[i] = bot[i-1] + int64(grid[1][i])
//     }
//     var res int64
//     res = math.MaxInt64
//     for i:= 0 ; i < n ; i++{
//         var t, b int64
//         t = top[n-1] - top[i]
//         b = 0
//         if i > 0{
//             b = bot[i-1]
//         }
//         a := max(t, b)
//         res = min(res, a)
//     }
//     return res
// }

// func min(a, b int64) int64{
//     if a > b{
//         return b
//     }
//     return a
// }

// func max(a, b int64)int64{
//     if a > b{
//         return a
//     }
//     return b
// }

func gridGame(gd [][]int) int64 {
	n := len(gd[0])
	var tT, tB int64
	for i := 0; i < n; i++ {
		tT += int64(gd[0][i])
		tB += int64(gd[1][i])
	}
	m := int64(math.MaxInt64)
	var cB int64
	for i := 0; i < n; i++ {
		tT -= int64(gd[0][i])
		s := x(tT, cB)
		m = y(m, s)
		cB += int64(gd[1][i])
	}
	return m
}

func x(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func y(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}