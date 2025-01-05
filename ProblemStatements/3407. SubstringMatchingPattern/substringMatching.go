func hasMatch(s string, p string) bool {
	h := strings.Split(p, "*")
	if len(h) != 2 {
		return false
	}
	temp := strings.ReplaceAll(p, "*", "")
	if strings.Contains(s, temp) {
		return true
	}
	l, r := h[0], h[1]
	if l == "" && r == "" {
		return true
	}
	if l == "" {
		return strings.HasSuffix(s, r)
	}
	if r == "" {
		return strings.HasPrefix(s, l)
	}
	for i := 0; i <= len(s)-len(l)-len(r); i++ {
		if strings.HasPrefix(s[i:], l) && strings.Contains(s[i+len(l):], r) {
			return true
		}
	}
	return false
}