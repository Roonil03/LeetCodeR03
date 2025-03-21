func maxPalindromes(s string, k int) int {
	// n := len(s)
	// dp := make([][]bool, n)
	// for i := range dp{
	//     dp[i] = make([]bool, n)
	//     dp[i][i] = true
	// }
	// for l := 2; l <= n; l++ {
	//     for i := 0; i <= n-l; i++ {
	//         j := i + l - 1
	//         if l == 2 {
	//             dp[i][j] = s[i] == s[j]
	//         } else {
	//             dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
	//         }
	//     }
	// }
	// d2 := make([]int, n+1)
	// for i := 1; i <= n; i++ {
	//     d2[i] = d2[i-1]
	//     for j := 0; j < i; j++ {
	//         if dp[j][i-1] && i-j >= k {
	//             d2[i] = max(d2[i], d2[j]+1)
	//         }
	//     }
	// }
	// return d2[n]
	if s == "" {
		return 0
	}
	for pair := range min(2, len(s)-k+1) {
		i, j := 0, k+pair-1
		for i <= j && s[i] == s[j] {
			i, j = i+1, j-1
		}
		if i < j {
			continue
		}
		return 1 + maxPalindromes(s[k+pair:], k)
	}
	return maxPalindromes(s[1:], k)
}