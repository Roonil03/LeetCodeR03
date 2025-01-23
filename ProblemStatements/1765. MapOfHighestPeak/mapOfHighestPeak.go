// func highestPeak(isWater [][]int) [][]int {
//     m, n := len(isWater), len(isWater[0]);
//     res := make([][]int, m)
//     for i := range res{
//         res[i] = make([] int, n)
//         for j := range res[i]{
//             res[i][j] = -1
//         }
//     }
//     q := list.New()
//     for i:= 0 ; i < m; i++{
//         for j := 0 ; j < n;j++{
//             if isWater[i][j] == 1{
//                 res[i][j] = 0
//                 q.PushBack([2]int{i, j})
//             }
//         }
//     }
//     for q.Len() > 0{
//         elem := q.Front()
//         q.Remove(elem)
//         c := elem.Value.([2]int)
//         cx, cy := c[0], c[1]
//         dir := [][2]int{
// 		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
// 	    }
//         for _, d := range dir{
//             nx, ny := cx + d[0] ,cy + d[1]
//             if nx >= 0 && nx < m && ny >= 0 && ny < n && res[nx][ny] == -1 {
// 				res[nx][ny] = res[cx][cy] + 1
// 				q.PushBack([2]int{nx, ny})
// 			}
//         }
//     }
//     return res
// }

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	dir := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	q := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = 0
				vis[i][j] = true
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		cx, cy := cell[0], cell[1]

		for _, d := range dir {
			nx, ny := cx+d[0], cy+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] {
				vis[nx][ny] = true
				res[nx][ny] = res[cx][cy] + 1
				q = append(q, [2]int{nx, ny})
			}
		}
	}

	return res
}
