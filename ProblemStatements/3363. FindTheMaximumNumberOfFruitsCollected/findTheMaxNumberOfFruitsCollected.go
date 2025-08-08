func maxCollectedFruits(fruits [][]int) int {
	// n := len(fruits)
	// type s struct{
	//     i1, j1, i2, j2, i3, j3 int
	// }
	// cur := map[[6]int]int{}
	// cur[[6]int{0, 0, 0, n - 1, n - 1, 0}] = fruits[0][0] + fruits[0][n-1] + fruits[n - 1][0]
	// sum := func(i1, j1, i2, j2, i3, j3 int) int{
	//     set := map[[2]int]bool{}
	//     pts := [][2]int{{i1, j1}, {i2, j2}, {i3, j3}}
	//     tot := 0
	//     for _, p := range pts {
	//         if !set[p] {
	//             tot += fruits[p[0]][p[1]]
	//             set[p] = true
	//         }
	//     }
	//     return tot
	// }
	// for k := 1; k<= n; k++{
	//     next := map[[6]int] int{}
	//     for pos, val := range cur{
	//         i1, j1, i2, j2, i3, j3 := pos[0], pos[1], pos[2], pos[3], pos[4], pos[5]
	//         m1 := [][2]int{}
	//         for _, d := range [][2]int{{1, 0}, {0, 1}, {1, 1}}{
	//             ni, nj := i1 + d[0], j1 + d[1]
	//             if ni < n && nj < n && ni + nj == k{
	//                 m1 = append(m1, [2]int{ni, nj})
	//             }
	//         }
	//         m2 := [][2]int{}
	//         for _, d := range [][2]int{{1, 0}, {1, -1}, {1, 1}}{
	//             ni, nj := i2 + d[0], j2 + d[0]
	//             if ni < n && ni >= 0 && nj < n && nj >= 0 && ni + (n - 1 - nj) == k{
	//                 m2 = append(m2, [2]int{ni, nj})
	//             }
	//         }
	//         m3:= [][2]int{}
	//         for _, d := range [][2]int{{-1, 1}, {0, 1}, {1, 1}}{
	//             ni, nj := i3+d[0], j3+d[1]
	//             if ni < n && ni >= 0 && nj < n && ni >= 0 && (n-1-ni)+nj == k {
	//                 m3 = append(m3, [2]int{ni, nj})
	//             }
	//         }
	//         for _, i := range m1{
	//             for _, j := range m2{
	//                 for _, k := range m3{
	//                     //tot := val + sum(i[0], i[1], j[0], j[1], k[0], k[1]) - sum(i1, j1, i2, j2, i3, j3)
	//                     key := [6]int{i[0], i[1], j[0], j[1], k[0], k[1]}
	//                     add := sum(i[0], i[1], j[0], j[1], k[0], k[1])
	//                     if nv := next[key]; nv < val + add{
	//                         next[key] = val + add
	//                     }
	//                 }
	//             }
	//         }
	//     }
	//     cur = next
	// }
	// res := 0
	// key := [6]int{n - 1, n - 1, n - 1, n - 1, n - 1, n - 1}
	// if v, ok := cur[key]; ok{
	//     res = v
	// } else{
	//     for k, v := range cur{
	//         vl := true
	//         for i := 0; i < 6; i += 2{
	//             if k[i] != n - 1 || k[i + 1] != n - 1{
	//                 vl = false
	//                 break
	//             }
	//         }
	//         if vl && v > res{
	//             res = v
	//         }
	//     }
	// }
	// return res

	// n := len(fruits)
	// cur := map[[6]int]int{}
	// cur[[6]int{0, 0, 0, n - 1, n - 1, 0}] = fruits[0][0] + fruits[0][n-1] + fruits[n-1][0]
	// sum := func(i1, j1, i2, j2, i3, j3 int) int {
	// 	set := map[[2]int]bool{}
	// 	pts := [][2]int{{i1, j1}, {i2, j2}, {i3, j3}}
	// 	tot := 0
	// 	for _, p := range pts {
	// 		if !set[p] {
	// 			tot += fruits[p[0]][p[1]]
	// 			set[p] = true
	// 		}
	// 	}
	// 	return tot
	// }
	// for k := 1; k < n; k++ {
	// 	next := map[[6]int]int{}
	// 	for pos, val := range cur {
	// 		i1, j1, i2, j2, i3, j3 := pos[0], pos[1], pos[2], pos[3], pos[4], pos[5]
	// 		m1 := [][2]int{}
	// 		for _, d := range [][2]int{{1, 0}, {0, 1}, {1, 1}} {
	// 			ni, nj := i1+d[0], j1+d[1]
	// 			if ni < n && nj < n && ni+nj == k {
	// 				m1 = append(m1, [2]int{ni, nj})
	// 			}
	// 		}
	// 		m2 := [][2]int{}
	// 		for _, d := range [][2]int{{1, 0}, {1, -1}, {1, 1}} {
	// 			ni, nj := i2+d[0], j2+d[1]
	// 			if ni < n && ni >= 0 && nj < n && nj >= 0 && ni+(n-1-nj) == k {
	// 				m2 = append(m2, [2]int{ni, nj})
	// 			}
	// 		}
	// 		m3 := [][2]int{}
	// 		for _, d := range [][2]int{{-1, 1}, {0, 1}, {1, 1}} {
	// 			ni, nj := i3+d[0], j3+d[1]
	// 			if ni < n && ni >= 0 && nj < n && nj >= 0 && (n-1-ni)+nj == k {
	// 				m3 = append(m3, [2]int{ni, nj})
	// 			}
	// 		}
	// 		for _, a := range m1 {
	// 			for _, b := range m2 {
	// 				for _, c := range m3 {
	// 					key := [6]int{a[0], a[1], b[0], b[1], c[0], c[1]}
	// 					sc := sum(a[0], a[1], b[0], b[1], c[0], c[1])
	// 					if nv := next[key]; nv < val+sc {
	// 						next[key] = val + sc
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	cur = next
	// }
	// res := 0
	// for k, v := range cur {
	// 	ok := true
	// 	for i := 0; i < 6; i += 2 {
	// 		if k[i] != n-1 || k[i+1] != n-1 {
	// 			ok = false
	// 			break
	// 		}
	// 	}
	// 	if ok && v > res {
	// 		res = v
	// 	}
	// }
	// return res

	n := len(fruits)
	total := 0
	for i := 0; i < n; i++ {
		total += fruits[i][i]
	}
	rightPath := make([]int, 3)
	rightPath[0] = fruits[0][n-1]
	bottomPath := make([]int, 3)
	bottomPath[0] = fruits[n-1][0]
	window := 2
	for step := 1; step < n-1; step++ {
		newRight := make([]int, window+2)
		newBottom := make([]int, window+2)
		for dist := 0; dist < window; dist++ {
			left, mid, right := 0, rightPath[dist], 0
			if dist-1 >= 0 {
				left = rightPath[dist-1]
			}
			if dist+1 < len(rightPath) {
				right = rightPath[dist+1]
			}
			newRight[dist] = max(left, max(mid, right)) + fruits[step][n-1-dist]
			left, mid, right = 0, bottomPath[dist], 0
			if dist-1 >= 0 {
				left = bottomPath[dist-1]
			}
			if dist+1 < len(bottomPath) {
				right = bottomPath[dist+1]
			}
			newBottom[dist] = max(left, max(mid, right)) + fruits[n-1-dist][step]
		}
		rightPath = newRight
		bottomPath = newBottom
		if window-n+4+step <= 1 {
			window++
		} else if window-n+3+step > 1 {
			window--
		}
	}
	return total + rightPath[0] + bottomPath[0]
}