package main

func finalPrices(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	s := []int{}
	for i := 0; i < n; i++ {
		for len(s) > 0 && prices[s[len(s)-1]] >= prices[i] {
			id := s[len(s)-1]
			s = s[:len(s)-1]
			res[id] = prices[id] - prices[i]
		}
		s = append(s, i)
	}
	for _, id := range s {
		res[id] = prices[id]
	}
	return res
}
