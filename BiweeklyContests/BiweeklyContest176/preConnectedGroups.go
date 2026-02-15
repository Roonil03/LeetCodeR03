func prefixConnected(words []string, k int) int {
	mp := make(map[string]int)
	for _, i := range words {
		if len(i) >= k {
			mp[(i[:k])]++
		}
	}
	res := 0
	for _, i := range mp {
		if i >= 2 {
			res++
		}
	}
	return res
}