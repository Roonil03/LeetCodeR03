func maxVowels(s string, k int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	res, count := 0, 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if i >= k && isVowel(s[i-k]) {
			count--
		}
		if count > res {
			res = count
		}
	}
	return res
}