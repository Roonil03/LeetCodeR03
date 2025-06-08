func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	chars1 := make([]bool, 26)
	chars2 := make([]bool, 26)
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
		chars1[word1[i]-'a'] = true
		chars2[word2[i]-'a'] = true
	}
	for i := 0; i < 26; i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}