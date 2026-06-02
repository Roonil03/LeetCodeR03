func maximumSaleItems(items [][]int, budget int) int {
	n := len(items)
	f := make([]int, n)
	temp := math.MaxInt
	for i := range n {
		temp = min(items[i][1], temp)
		count := 0
		for j := range n {
			if items[j][0]%items[i][0] == 0 {
				count++
			}
		}
		f[i] = count
	}
	dp := make([]int, budget+1)
	for i := 0; i <= budget; i++ {
		dp[i] = i / temp
	}
	for i := range n {
		a, b := items[i][1], f[i]
		for j := budget; j >= a; j-- {
			dp[j] = max(dp[j], dp[j-a]+b)
		}
	}
	res := 0
	for i := 0; i <= budget; i++ {
		res = max(res, dp[i]+(budget-i)/temp)
	}
	return res
}