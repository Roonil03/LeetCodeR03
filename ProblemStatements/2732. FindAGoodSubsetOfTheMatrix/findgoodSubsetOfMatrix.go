func goodSubsetofBinaryMatrix(grid [][]int) []int {
	// n, m := len(grid), len(grid[0])
	// for i := 0; i < n; i++{
	//     z := true
	//     for j := 0; j < m; j++{
	//         if grid[i][j] == 1{
	//             z = false
	//             break
	//         }
	//     }
	//     if z{
	//         return []int{i}
	//     }
	// }
	// for i := 0; i < n; i++{
	//     for j := i + 1; j < n; j++{
	//         comp := true
	//         for k := 0; k < m; k++{
	//             if grid[i][k] == 1 && grid[j][k] == 1{
	//                 comp = false
	//                 break
	//             }
	//         }
	//         if comp{
	//             return []int{i, j}
	//         }
	//     }
	// }
	// return []int{}
	n, m := len(grid), len(grid[0])
	N := 1 << m
	s := make([]int, N)
	for i := range s {
		s[i] = -1
	}
	encode := func(row []int) int {
		var x int
		for i := 0; i < len(row); i++ {
			if row[i] != 0 {
				x |= 1 << i
			}
		}
		return x
	}
	for i := n - 1; i >= 0; i-- {
		s[encode(grid[i])] = i
	}
	if s[0] != -1 {
		return []int{s[0]}
	}
	for a := 1; a < N; a++ {
		if s[a] == -1 {
			continue
		}
		for b := a + 1; b < N; b++ {
			if s[b] == -1 {
				continue
			}
			if (a & b) == 0 {
				ra, rb := s[a], s[b]
				if ra < rb {
					return []int{ra, rb}
				} else {
					return []int{rb, ra}
				}
			}
		}
	}
	return []int{}
}