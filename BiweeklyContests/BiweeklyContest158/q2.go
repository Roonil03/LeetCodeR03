func maximumProfit(prices []int, k int) int64 {
	// n := len(prices)
	// if k >= n{
	//     res := int64(0)
	//     for i:= 1; i < n; i++{
	//         res += int64(max(prices[i] - prices[i-1], 0))
	//         res += int64(max(prices[i-1] - prices[i], 0))
	//     }
	//     return res
	// }
	// dp := make([][2]int, k+1)
	// for t := 1; t<= k; t++{
	//     dp[t][0], dp[t][1] = -1 << 31, -1<<31
	// }
	// for i:= 0; i < n; i++{
	//     for t := k; t >= 1; t--{
	//         dp[t][0] = max(dp[t][0], dp[t-1][1]+prices[i])
	//         dp[t][1] = max(max(dp[t][1], dp[t-1][0] - prices[i]), dp[t-1][0]+prices[i])
	//     }
	//     dp[0][0], dp[0][1] = 0, -1<<31
	// }
	// res := int64(0)
	// for t :=  0; t <= k; t++{
	//     res = max64(res, int64(dp[t][0]))
	// }
	// return res
	//
	dp := make([][3]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = [3]int{-1 << 31, -1 << 31, -1 << 31}
	}
	dp[0][0] = 0
	for _, p := range prices {
		o := make([][3]int, k+1)
		for i := 0; i <= k; i++ {
			o[i] = [3]int{-1 << 31, -1 << 31, -1 << 31}
		}
		for i := 0; i <= k; i++ {
			if dp[i][0] != -1<<31 {
				if o[i][0] < dp[i][0] {
					o[i][0] = dp[i][0]
				}
				if i < k {
					o[i][1] = max(o[i][1], dp[i][0]-p)
					o[i][2] = max(o[i][2], dp[i][0]+p)
				}
			}
			if dp[i][1] != -1<<31 {
				if i+1 <= k && o[i+1][0] < dp[i][1]+p {
					o[i+1][0] = dp[i][1] + p
				}
				if o[i][1] < dp[i][1] {
					o[i][1] = dp[i][1]
				}
			}
			if dp[i][2] != -1<<31 {
				if i+1 <= k && o[i+1][0] < dp[i][2]-p {
					o[i+1][0] = dp[i][2] - p
				}
				if o[i][2] < dp[i][2] {
					o[i][2] = dp[i][2]
				}
			}
		}
		dp = o
	}
	res := int64(0)
	for i := 0; i <= k; i++ {
		if int64(dp[i][0]) > res {
			res = int64(dp[i][0])
		}
	}
	return res
}

// func max64(a, b int64)int64{
//     if a > b{
//         return a
//     }
//     return b
// }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}