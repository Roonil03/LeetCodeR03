func longestDecomposition(text string) int {
	n := len(text)
	for i := 0; i < n/2; i++ {
		if text[:i+1] == text[n-1-i:] {
			return 2 + longestDecomposition(text[i+1:n-i-1])
		}
	}
	if n == 0 {
		return 0
	}
	return 1
}