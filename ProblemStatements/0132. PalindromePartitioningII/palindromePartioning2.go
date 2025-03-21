func minCut(s string) int {
	// n := len(s)
	// if n <= 1{
	//     return 0
	// }
	// dp := make([]int, n)
	// pal := make([][]bool, n)
	// for i:= range pal{
	//     pal[i] = make([]bool, n)
	//     pal[i][i] = true
	// }
	// for a := 2; a <= n; a++{
	//     for b := 0; b <= n-a; b++ {
	//         c := b + a - 1
	//         if a == 2 {
	//             pal[b][c] = s[b] == s[c]
	//         } else {
	//             pal[b][c] = s[b] == s[c] && pal[b+1][c-1]
	//         }
	//     }
	// }
	// for i := 0; i < n; i++ {
	//     if pal[0][i] {
	//         dp[i] = 0
	//         continue
	//     }
	//     dp[i] = i
	//     for j := 0; j < i; j++ {
	//         if pal[j+1][i] {
	//             dp[i] = min(dp[i], dp[j] + 1)
	//         }
	//     }
	// }
	// return dp[n-1]
	N := len(s)
	minCut := make([]int, N) // minCut[l] == minumum cut for s[:l+1]
	for i := 0; i < N; i++ {
		minCut[i] = i // length of l always has a cut l-1
	}
	update := func(l, r int) {
		for l >= 0 && r < N && s[l] == s[r] {
			if l == 0 {
				minCut[r] = 0 // s[:r+1] is a palindrome already
			} else {
				// s[l:r+1] is a palindrome, try this cut
				minCut[r] = min(minCut[r], minCut[l-1]+1)
			}
			l, r = l-1, r+1
		}
	}
	for c := 1; c < N; c++ {
		// use c as center
		update(c, c)
		update(c-1, c)
	}
	return minCut[N-1]
}