func maximumProduct(nums []int, m int) int64 {
	n := len(nums)
	if n < m {
		return 0
	}
	// dp := make([][]int64, n)
	// for i := range dp{
	//     dp[i] = make([]int64, m+1)
	//     for j := range dp[i]{
	//         dp[i][j] = int64(math.MinInt32)
	//     }
	// }
	// for i := 0; i < n; i++{
	//     dp[i][1] = int64(nums[i])
	// }
	// res := int64(math.MinInt32)
	// for k := 2; k <= m; k++{
	//     for i := k - 1; i < n; i++{
	//         for p := k - 2; p < i; p++{
	//             if dp[p][k-1] != int64(math.MinInt32){
	//                 if k == m{
	//                     a := dp[p][1]
	//                     prod := a * int64(nums[i])
	//                     res = max(res, prod)
	//                 } else{
	//                     dp[i][k] = max(dp[i][k], int64(nums[i]))
	//                 }
	//             }
	//         }
	//     }
	// }
	// if m == 1{
	//     prod := int64(nums[0] * nums[0])
	//     for i := 1; i < n; i++{
	//         p := int64(nums[i] * nums[i])
	//         if p > prod{
	//             prod = p
	//         }
	//     }
	//     return prod
	// }
	// res := int64(-1e11)
	// for i := 0; i <= n - m; i++{
	//     for j := i + m - 1;j < n; j++{
	//         p := int64(nums[i]) * int64(nums[j])
	//         if p > res{
	//             res = p
	//         }
	//     }
	// }
	if m == 1 {
		res := int64(nums[0] * nums[0])
		for i := 1; i < n; i++ {
			p := int64(nums[i] * nums[i])
			if p > res {
				res = p
			}
		}
		return res
	}
	x := make([]int, n)
	y := make([]int, n)
	x[n-1] = nums[n-1]
	y[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > x[i+1] {
			x[i] = nums[i]
		} else {
			x[i] = x[i+1]
		}
		if nums[i] < y[i+1] {
			y[i] = nums[i]
		} else {
			y[i] = y[i+1]
		}
	}
	res := int64(-1e18)
	for i := 0; i <= n-m; i++ {
		k := i + m - 1
		a := int64(nums[i])
		var p int64
		if nums[i] > 0 {
			p = a * int64(x[k])
		} else if nums[i] < 0 {
			p = a * int64(y[k])
		} else {
			p = 0
		}
		if p > res {
			res = p
		}
	}
	return res
}