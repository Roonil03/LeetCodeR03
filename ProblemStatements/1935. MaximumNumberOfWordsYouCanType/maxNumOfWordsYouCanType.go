func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		broken[brokenLetters[i]] = true
	}
	words := strings.Split(text, " ")
	count := 0
	for _, word := range words {
		canType := true
		for i := 0; i < len(word); i++ {
			if broken[word[i]] {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}