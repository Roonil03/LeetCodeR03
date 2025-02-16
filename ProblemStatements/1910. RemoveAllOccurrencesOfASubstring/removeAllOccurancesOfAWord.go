func removeOccurrences(s string, part string) string {
	pl := len(part)
	for {
		i := strings.Index(s, part)
		if i == -1 {
			break
		}
		s = s[:i] + s[i+pl:]
	}
	return s
}