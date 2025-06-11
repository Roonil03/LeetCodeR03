func maxDifference(s string, k int) int {
	//     res, n := -1, len(s)
	//     for i:=k ;i <=n;i++{
	//         freq := make([]int, 26)
	//         for j := 0; j < i; j++{
	//             freq[s[j] - '0']++
	//         }
	//         res = max(res, h1(freq))
	//         for j := i; j < n; j++{
	//             freq[s[j]-'0']++
	//             freq[s[j-i] - '0']--
	//             res = max(res, h1(freq))
	//         }
	//     }
	//     return res
	// }

	// func h1(freq []int)int{
	//     res := -1
	//     for i:= 0; i < 26; i++{
	//         if freq[i] > 0 && freq[i] % 2 == 1{
	//             for j := 0; j < 26; j++{
	//                 if freq[j] > 0 && freq[j] % 2 == 0{
	//                     d := freq[i] - freq[j]
	//                     res = max(res, d)
	//                 }
	//             }
	//         }
	//     }
	//     return res
	// }
	n := len(s)
	vals := make([]int, n)
	for i := range s {
		vals[i] = int(s[i] - '0')
	}
	res := -1000000000
	d := make([]int, n+1)
	pa := make([]int, n+1)
	cb := make([]int, n+1)
	cb_list := make([][]int, 4)
	dm := make([][]int, 4)
	for t := 0; t < 4; t++ {
		cb_list[t] = make([]int, n+1)
		dm[t] = make([]int, n+1)
	}
	ptr := [4]int{}
	sz := [4]int{}
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			if a == b {
				continue
			}
			d[0], pa[0], cb[0] = 0, 0, 0
			for i := 1; i <= n; i++ {
				da := 0
				if vals[i-1] == a {
					da = 1
				}
				db := 0
				if vals[i-1] == b {
					db = 1
				}
				d[i] = d[i-1] + da - db
				pa[i] = pa[i-1] ^ da
				cb[i] = cb[i-1] + db
			}
			for t := 0; t < 4; t++ {
				ptr[t], sz[t] = 0, 0
			}
			for j := k; j <= n; j++ {
				i := j - k
				idx := (pa[i] << 1) | (cb[i] & 1)
				di := d[i]
				szi := sz[idx]
				if szi == 0 {
					dm[idx][0] = di
				} else {
					prev := dm[idx][szi-1]
					if di < prev {
						dm[idx][szi] = di
					} else {
						dm[idx][szi] = prev
					}
				}
				cb_list[idx][szi] = cb[i]
				sz[idx] = szi + 1
				tidx := ((pa[j] ^ 1) << 1) | (cb[j] & 1)
				T := cb[j] - 2
				p := ptr[tidx]
				for p < sz[tidx] && cb_list[tidx][p] <= T {
					p++
				}
				ptr[tidx] = p
				if p > 0 {
					diff := d[j] - dm[tidx][p-1]
					if diff > res {
						res = diff
					}
				}
			}
		}
	}
	return res
}
