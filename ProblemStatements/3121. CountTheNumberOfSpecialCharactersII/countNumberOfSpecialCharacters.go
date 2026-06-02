func numberOfSpecialChars(word string) int {
	up, low := make([]int, 26), make([]int, 26)
	for i := range 26 {
		up[i] = -1
		low[i] = -1
	}
	for i := range word {
		ch := word[i]
		if ch >= 'a' && ch <= 'z' {
			low[ch-'a'] = i
		}
		if ch >= 'A' && ch <= 'Z' && up[ch-'A'] == -1 {
			up[ch-'A'] = i
		}
	}
	res := 0
	for i := range 26 {
		if low[i] < up[i] && low[i] != -1 && up[i] != -1 {
			res++
		}
	}
	return res
}