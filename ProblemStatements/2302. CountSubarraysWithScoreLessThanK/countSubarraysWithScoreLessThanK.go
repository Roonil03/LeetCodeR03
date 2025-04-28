func countSubarrays(nums []int, k int64) int64 {
	//     n, ps := len(nums), make([]int64, len(nums))
	//     ps[0] = int64(nums[0])
	//     for i:= 1; i < n; i++{
	//         ps[i] = ps[i-1] + int64(nums[i])
	//     }
	//     var res int64 = 0
	//     for i:=0;i < n; i++{
	//         r := binarySearch(i, n-1, ps, k)
	//         if r != -1{
	//             res += int64(r - i + 1)
	//         }
	//     }
	//     return res-1
	// }

	// func binarySearch(st, end int, ps []int64, k int64)int{
	//     r := -1
	//     pos := st
	//     for st <= end{
	//         m := (st + end) / 2
	//         l := int64(m - st + 1)
	//         sum := ps[m] - func () int64{
	//             if pos == 0{
	//                 return 0
	//             }
	//             return ps[pos-1]
	//         } ()
	//         score := l * sum
	//         if score >= k{
	//             end = m - 1
	//         } else{
	//             r = m
	//             st = m + 1
	//         }
	//     }
	//     return r
	// }
	var s, res int64 = 0, 0
	j := 0
	for i := 0; i < len(nums); i++ {
		s += int64(nums[i])
		for s*int64(i-j+1) >= k {
			s -= int64(nums[j])
			j++
		}
		res += int64(i - j + 1)
	}
	return res
}