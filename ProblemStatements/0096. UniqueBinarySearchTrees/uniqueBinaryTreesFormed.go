// func numTrees(n int) int {
//     var temp int64
//     temp = fact(2*n) / fact(n) / fact(n) / int64(n+1)
//     return int(temp)
// }

// func fact(n int)int64{
//     var res int64
//     res = 1
//     for i := 2; i <=n; i++{
//         res *= int64(i)
//     }
//     return res
// }

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}