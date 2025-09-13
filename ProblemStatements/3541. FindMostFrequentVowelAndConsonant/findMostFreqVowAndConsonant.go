func maxFreqSum(s string) int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}
	v, c := 0, 0
	for i := range freq {
		ch := i + 'a'
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			v = max(v, freq[i])
		} else {
			c = max(c, freq[i])
		}
	}
	return v + c
}