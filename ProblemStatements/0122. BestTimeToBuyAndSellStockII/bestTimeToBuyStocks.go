func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max((prices[i] - prices[i-1]), 0)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}