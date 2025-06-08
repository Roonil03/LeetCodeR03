func removeStars(s string) string {
	x := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			x = append(x, s[i])
		} else if len(x) > 0 {
			x = x[:len(x)-1]
		}
	}
	return string(x)
}