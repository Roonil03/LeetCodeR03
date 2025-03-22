func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	// dp := make([][][]int64, n+1)
	// for i := range dp {
	//     dp[i] = make([][]int64, k+1)
	//     for j := range dp[i] {
	//         dp[i][j] = make([]int64, 2)
	//         for m := range dp[i][j] {
	//             dp[i][j][m] = -1e15
	//         }
	//     }
	// }
	// dp[0][0][0] = 0
	// mul := make([]int64, k)
	// for i := 0; i < k; i++ {
	//     if i%2 == 0 {
	//         mul[i] = int64(k - i)
	//     } else {
	//         mul[i] = int64(-k + i)
	//     }
	// }
	// for i := 0; i < n; i++ {
	//     for j := 0; j <= k; j++ {
	//         for m := 0; m < 2; m++ {
	//             curr := dp[i][j][m]
	//             if curr == -1e15 {
	//                 continue
	//             }
	//             dp[i+1][j][0] = max(dp[i+1][j][0], curr)
	//             if m == 0 && j < k {
	//                 dp[i+1][j+1][1] = max(dp[i+1][j+1][1], curr + mul[j]*int64(nums[i]))
	//             }
	//             if m == 1 {
	//                 next := curr + mul[j-1]*int64(nums[i])
	//                 dp[i+1][j][1] = max(dp[i+1][j][1], next)
	//             }
	//         }
	//     }
	// }
	// if dp[n][k][0] > dp[n][k][1] {
	//     return dp[n][k][0]
	// }
	// return dp[n][k][1]
	dp := make([]int64, n+1)
	for j := 0; j < k; j++ {
		mul := int64(k - j)
		if j%2 == 1 {
			mul = -mul
		}
		tmp := make([]int64, n+1)
		tmp[0] = math.MinInt64 / 2
		for i := 1; i <= n; i++ {
			tmp[i] = max(tmp[i-1], dp[i-1]) + int64(nums[i-1])*mul
		}
		copy(dp, tmp)
		for i := 1; i <= n; i++ {
			dp[i] = max(dp[i], dp[i-1])
		}
	}
	return dp[n]
}