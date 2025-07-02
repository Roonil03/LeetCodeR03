func possibleStringCount(word string, k int) int {
	const mod = int(1e9 + 7)
	// n := len(word)
	// dp := make([]map[int]int, n + 1)
	// dp[0] = make(map[int]int)
	// dp[0][0] = 1
	// type group struct{
	//     s, e int
	// }
	// g := []group{}
	// i := 0
	// for i < n{
	//     j := i
	//     for j+ 1< n && word[j + 1] == word[i]{
	//         j++
	//     }
	//     g = append(g, group{i, j})
	//     i = j + 1
	// }
	// for i := 0; i < len(g); i++{
	//     gr := g[i]
	//     l := gr.e -gr.s + 1
	//     DP := make(map[int]int)
	//     for o , c := range dp[gr.s]{
	//         for keep := 1; keep <= l; keep++{
	//             nl := o + keep
	//             DP[nl] = (DP[nl] + c) % mod
	//         }
	//     }
	//     dp[gr.e + 1] = DP
	// }
	// res := 0
	// for i, c := range dp[n]{
	//     if i >= k{
	//         res = (res + c) % mod
	//     }
	// }
	// return res
	if len(word) == k {
		return 1
	}
	if len(word) < k {
		return 0
	}
	clumsy_arr := []int{1}
	total_char := 1
	res := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			clumsy_arr[len(clumsy_arr)-1]++
		} else {
			total_char++
			if clumsy_arr[len(clumsy_arr)-1] > 1 {
				res = (res * (clumsy_arr[len(clumsy_arr)-1])) % mod
				clumsy_arr[len(clumsy_arr)-1]--
				clumsy_arr = append(clumsy_arr, 1)
			}
		}
	}
	res = (res * (clumsy_arr[len(clumsy_arr)-1])) % mod
	if clumsy_arr[len(clumsy_arr)-1] > 1 {
		clumsy_arr[len(clumsy_arr)-1]--
	} else {
		clumsy_arr = clumsy_arr[:len(clumsy_arr)-1]
	}
	if len(clumsy_arr) == 0 {
		return 1
	}
	if total_char >= k {
		return res
	}
	total := k - total_char
	dp := make([]int, total)
	prefix := make([]int, total+1)
	for i := range dp {
		if clumsy_arr[len(clumsy_arr)-1] >= i {
			dp[i] = 1
		}
		prefix[i+1] = prefix[i] + dp[i]
	}
	for j := len(clumsy_arr) - 2; j >= 0; j-- {
		prefixN := make([]int, total+1)
		for i := range dp {
			dp[i] = prefix[i+1] - prefix[max(0, i-clumsy_arr[j])]
			prefixN[i+1] = (prefixN[i] + dp[i]) % mod
		}
		prefix = prefixN
	}
	return (res - prefix[total] + mod) % mod
}