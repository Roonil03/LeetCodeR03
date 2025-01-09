func prefixCount(words []string, pref string) int {
	res := 0
	n := len(words)
	for i := 0; i < n; i++ {
		if strings.HasPrefix(words[i], pref) {
			res++
		}
	}
	return res
}