// func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int {
//     b := 225
//     n2 := len(nums2)
//     n := (n2 + b - 1) / b
//     freq := make([]map[int]int, n)
//     lazy := make([]int, n)
//     for i := range n{
//         freq[i] = make(map[int]int)
//     }
//     for i, v := range nums2{
//         freq[i / b][v]++
//     }
//     var res []int
//     for _, q := range queries{
//         if q[0] == 1{
//             l, r, v := q[1], q[2], q[3]
//             bl, br := l/b, r/b
//             if bl == br{

//             }
//         } else{
//             count := 0
//             for _, v := range nums2{
//                 count += freq[q[1]- v]
//             }
//             res = append(res, count)
//         }
//     }
//     return res
// }

func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int {
	b := 225
	n2 := len(nums2)
	n := (n2 + b - 1) / b
	freq := make([]map[int]int, n)
	lazy := make([]int, n)
	for i := range n {
		freq[i] = make(map[int]int)
	}
	for i, v := range nums2 {
		freq[i/b][v]++
	}
	var res []int
	for _, q := range queries {
		if q[0] == 1 {
			l, r, val := q[1], q[2], q[3]
			bl, br := l/b, r/b
			if bl == br {
				for i := l; i <= r; i++ {
					freq[bl][nums2[i]]--
					nums2[i] += val
					freq[bl][nums2[i]]++
				}
			} else {
				for i := l; i < (bl+1)*b; i++ {
					freq[bl][nums2[i]]--
					nums2[i] += val
					freq[bl][nums2[i]]++
				}
				for i := bl + 1; i < br; i++ {
					lazy[i] += val
				}
				for i := br * b; i <= r; i++ {
					freq[br][nums2[i]]--
					nums2[i] += val
					freq[br][nums2[i]]++
				}
			}
		} else {
			tot := q[1]
			aa := 0
			for _, v := range nums1 {
				for i := range n {
					aa += freq[i][tot-v-lazy[i]]
				}
			}
			res = append(res, aa)
		}
	}
	return res
}