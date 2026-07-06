func canMakeSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	a, b := 0, 0
	for i := 0; i < len(t) && b < len(s); i++ {
		if s[b] == t[i] || a == b {
			b++
		}
		if s[a] == t[i] {
			a++
		}
	}
	return b == len(s)
}