func findWordsContaining(words []string, x byte) []int {
	res := []int{}
	for t, w := range words {
		for i := 0; i < len(w); i++ {
			if x == w[i] {
				res = append(res, t)
				break
			}
		}
	}
	return res
}