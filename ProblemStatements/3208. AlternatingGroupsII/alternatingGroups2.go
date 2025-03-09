func numberOfAlternatingGroups(colors []int, k int) int {
	res := 0
	// n := len(colors)
	// for i:= 0; i< len(colors); i++{
	//     a := true
	//     for j := i+1; j < i+k; j++{
	//         if colors[j%n] == colors[(j-1)%n]{
	//             a = !a
	//             break
	//         }
	//     }
	//     if a{
	//         res++
	//     }
	// }
	// a := true
	// for i:= 0; i < k; i++{
	//     if colors[i] == colors[(i+1)%n]{
	//         a = !a
	//         break
	//     }
	// }
	// if a{
	//     res++
	// }
	// for i:= 1; i <n;i++{
	//     if colors[(i-1)%n] == colors[(i+k)%n]{
	//         a = false
	//     }
	//     if a{
	//         res++
	//     }
	//     a = true
	//     for j := 0; j < k; j++ {
	// 		if colors[(i+j)%n] == colors[(i+j+1)%n] {
	// 			a = false
	// 			break
	// 		}
	// 	}
	// 	if a {
	// 		res++
	// 	}
	// }

	colors = append(colors, colors[:k-1]...)
	j := 0
	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i] == colors[(i-1)] {
			j = i
		}
		if i-j+1 >= k {
			res++
		}
	}
	return res
}