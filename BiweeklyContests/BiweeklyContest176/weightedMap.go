func mapWordWeights(words []string, weights []int) string {
	// res := "";
	var res strings.Builder
	res.Grow(len(words))
	for _, s := range words {
		g := 0
		for _, c := range s {
			g += weights[c-'a']
		}
		g %= 26
		h := 'a' + byte(25-g)
		res.WriteByte(h)
		// res = ('a' + byte(25-g)) + res
	}
	return res.String()
}