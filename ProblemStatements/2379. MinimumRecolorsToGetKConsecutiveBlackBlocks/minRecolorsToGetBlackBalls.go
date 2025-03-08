// func minimumRecolors(blocks string, k int) int {
//     n := len(blocks)
//     res := n
//     ops := 0
//     for i:=0;i<k;i++{
//         if blocks[i] == 'W'{
//             ops++
//         }
//     }
//     for i := k; i < n; i++{
//         if blocks[i] == 'W'{
//             ops++
//         }
//         if blocks[i-k] == 'W'{
//             ops--
//         }
//         res = min(res, ops)
//     }
//     return min(res, ops)
// }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumRecolors(blocks string, k int) int {
	w := 0
	res := math.MaxInt
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			w++
		}
		if i >= k-1 {
			res = min(res, w)
			if blocks[i-(k-1)] == 'W' {
				w--
			}
		}
	}
	return res
}