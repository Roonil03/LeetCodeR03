func countConsistentStrings(allowed string, words []string) int {
	allowedMask := 0
	for i := 0; i < len(allowed); i++ {
		allowedMask |= 1 << (allowed[i] - 'a')
	}
	count := 0
	for _, word := range words {
		wordMask := 0
		for i := 0; i < len(word); i++ {
			wordMask |= 1 << (word[i] - 'a')
		}
		if wordMask&allowedMask == wordMask {
			count++
		}
	}
	return count
}