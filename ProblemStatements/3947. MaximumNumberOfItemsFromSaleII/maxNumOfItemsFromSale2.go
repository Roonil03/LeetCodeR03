func maximumSaleItems(items [][]int, budget int) int {
	minPrice := math.MaxInt
	maxFactor := 0
	for _, item := range items {
		if item[1] < minPrice {
			minPrice = item[1]
		}
		if item[0] > maxFactor {
			maxFactor = item[0]
		}
	}
	count := make([]int, maxFactor+1)
	for _, item := range items {
		count[item[0]]++
	}
	multiples := make([]int, maxFactor+1)
	for f := 1; f <= maxFactor; f++ {
		if count[f] > 0 {
			for m := f; m <= maxFactor; m += f {
				multiples[f] += count[m]
			}
		}
	}
	var options []Option
	for _, item := range items {
		f := item[0]
		p := item[1]
		validPairs := multiples[f] - 1

		if validPairs > 0 && p < 2*minPrice {
			options = append(options, Option{price: p, count: validPairs})
		}
	}
	sort.Slice(options, func(i, j int) bool {
		return options[i].price < options[j].price
	})
	res := 0
	for _, opt := range options {
		if budget < opt.price {
			break
		}
		take := opt.count
		if budget/opt.price < take {
			take = budget / opt.price
		}
		res += take * 2
		budget -= take * opt.price
	}
	res += budget / minPrice
	return res
}

type Option struct {
	price int
	count int
}

// func maximumSaleItems(items [][]int, budget int) int {
//     n := len(items)
//     temp := math.MaxInt
//     for _, v := range items{
//         temp = min(v[1], temp)
//     }
//     dp := make([]int, budget + 1)
//     for i := range dp{
//         dp[i] = i / temp
//     }
//     for i := range n{
//         add := 0
//         for j := range n{
//             if i != j && items[j][0] % items[i][0] == 0{
//                 add++
//             }
//         }
//         cost := items[i][1]a
//         for j := 1; add > 0; j <<= 1{
//             if j > add{
//                 j = add
//             }
//             add -= j
//             for k := budget; k >= cost * j; k--{
//                 if v := dp[k - cost * j] + 2 * j; v > dp[k]{
//                     dp[k] = v
//                 }
//             }
//         }
//     }
//     return dp[budget]
// }