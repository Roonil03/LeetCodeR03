func longestPalindrome(words []string) int {
	freq := make(map[string]int)
	for _, ch := range words {
		freq[ch]++
	}
	res := 0
	mid := false
	for w, count := range freq {
		rev := string([]byte{w[1], w[0]})
		if w[0] == w[1] {
			p := count / 2
			res += p * 4
			if count%2 == 1 && !mid {
				res += 2
				mid = true
			}
		} else if w < rev {
			if v, ok := freq[rev]; ok {
				p := min(count, v)
				res += p * 4
			}
		}
	}
	return res
}