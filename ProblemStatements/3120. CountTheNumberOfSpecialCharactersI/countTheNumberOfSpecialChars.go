func numberOfSpecialChars(word string) int {
	up, low := make([]bool, 26), make([]bool, 26)
	for i := range word {
		ch := word[i]
		if ch >= 'A' && ch <= 'Z' {
			up[ch-'A'] = true
		}
		if ch >= 'a' && ch <= 'z' {
			low[ch-'a'] = true
		}
	}
	res := 0
	for i := range 26 {
		if up[i] && low[i] {
			res++
		}
	}
	return res
}