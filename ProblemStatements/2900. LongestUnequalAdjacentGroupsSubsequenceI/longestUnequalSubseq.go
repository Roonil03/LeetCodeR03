func getLongestSubsequence(words []string, groups []int) []string {
	res := []string{words[0]}
	for i := 1; i < len(groups); i++ {
		if groups[i] != groups[i-1] {
			res = append(res, words[i])
		}
	}
	return res
}