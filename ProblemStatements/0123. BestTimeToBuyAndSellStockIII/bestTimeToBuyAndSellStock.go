func maxProfit(prices []int) int {
	// if len(prices) <= 1{
	//     return 0
	// }
	// a := make([]int, len(prices))
	// m := prices[0]
	// for i := 1; i < len(prices); i++ {
	//     if prices[i] < m {
	//         m = prices[i]
	//     }
	//     a[i] = max(a[i-1], prices[i]-m)
	// }
	// m2 := prices[len(prices)-1]
	// b := 0
	// res := 0
	// for i := len(prices)-2; i >= 0; i-- {
	//     if prices[i] > m2 {
	//         m2 = prices[i]
	//     }
	//     b = max(b, m2-prices[i])
	//     res = max(res, a[i]+b)
	// }
	// return res
	hold1, sold1, hold2, sold2 := -100000, 0, -100000, 0
	for _, price := range prices {
		hold1, sold1, hold2, sold2 = max(hold1, -price), max(sold1, hold1+price), max(hold2, sold1-price), max(sold2, hold2+price)
	}
	return max(sold1, sold2)
}