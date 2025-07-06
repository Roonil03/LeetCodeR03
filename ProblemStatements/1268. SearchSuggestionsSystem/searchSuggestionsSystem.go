func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := [][]string{}
	pre := ""
	for _, ch := range searchWord {
		pre += string(ch)
		sug := []string{}
		i := sort.Search(len(products), func(i int) bool {
			return products[i] >= pre
		})
		for j := i; j < len(products) && len(sug) < 3; j++ {
			if strings.HasPrefix(products[j], pre) {
				sug = append(sug, products[j])
			} else {
				break
			}
		}
		res = append(res, sug)
	}
	return res
}