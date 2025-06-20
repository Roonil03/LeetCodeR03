// func maxDistance(s string, k int) int {
//     n := len(s)
//     type pt struct{
//         x, y int
//     }
//     dp := make([][]pt, n + 1)
//     for i:= 0; i <= n; i++{
//         dp[i] = make([]pt, k + 1)
//         for j := 0; j <= k; j++{
//             dp[i][j] = pt{0,0}
//         }
//     }
//     res := 0
//     vis := make([]map[pt]struct{}, n + 1)
//     for i:= range vis{
//         vis[i] = make(map[pt]struct{})
//     }
//     type state struct{
//         i, j, x, y int
//     }
//     q := []state{{0, 0, 0, 0}}
//     for len(q) > 0{
//         c := q[0]
//         q = q[1:]
//         if c.i == n{
//             continue
//         }
//         p := pt{c.x, c.y}
//         if _, ok := vis[c.i][p]; ok {
//             continue
//         }
//         vis[c.i][p] = struct{}{}
//         nx, ny := c.x, c.y
//         switch s[c.i] {
//         case 'N':
//             ny++
//         case 'S':
//             ny--
//         case 'E':
//             nx++
//         case 'W':
//             nx--
//         }
//         d := abs(nx) + abs(ny)
//         if d > res {
//             res = d
//         }
//         q = append(q, state{c.i + 1, c.j, nx, ny})
//         if c.j < k{
//             dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
//             for _, dir := range dirs {
//                 nx2 := c.x + dir[0]
//                 ny2 := c.y + dir[1]
//                 if (nx2 == nx && ny2 == ny) {
//                     continue
//                 }
//                 d2 := abs(nx2) + abs(ny2)
//                 if d2 > res {
//                     res = d2
//                 }
//                 q = append(q, state{c.i + 1, c.j + 1, nx2, ny2})
//             }
//         }
//     }
//     return res
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxDistance(s string, k int) int {
	n := 0
	e := 0
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'N':
			n++
		case 'S':
			n--
		case 'E':
			e++
		case 'W':
			e--
		}
		res = max(res, min(i+1, abs(n)+abs(e)+2*k))
	}
	return res
}