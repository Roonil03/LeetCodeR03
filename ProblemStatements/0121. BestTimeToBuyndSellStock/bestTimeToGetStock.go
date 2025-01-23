func maxProfit(prices []int) int {
	a, b := 0, 0
	for i := 1; i < len(prices); i++ {
		a += prices[i] - prices[i-1]
		a = max(0, a)
		b = max(a, b)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}