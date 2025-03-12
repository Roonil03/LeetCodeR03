func numDistinct(s string, t string) int {
	// m, n := len(s), len(t)
	// dp := make([][]int, m+1)
	// for i:= range dp{
	//     dp[i] = make([]int, n+1)
	// }
	// for i:=0;i<=m;i++{
	//     dp[i][0] = 1
	// }
	// for i:= 1; i <=m;i++{
	//     for j:= 1; j <= n;j++{
	//         if s[i-1] == t[j-1]{
	//             dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	//         } else{
	//             dp[i][j] = dp[i-1][j]
	//         }
	//     }
	// }
	// return dp[m][n]
	n, m := len(s), len(t)
	arr := make([]int, n)
	count, temp := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == t[0] {
			arr[i] = 1
		}
	}
	for i := 1; i < m; i++ {
		count = 0
		for j := 0; j < n; j++ {
			temp = arr[j]
			if s[j] == t[i] {
				arr[j] = count
			} else {
				arr[j] = 0
			}
			count += temp
		}
	}
	count = 0
	for j := 0; j < n; j++ {
		count += arr[j]
	}
	return count
}