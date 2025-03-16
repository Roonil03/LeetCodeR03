func freqAlphabets(s string) string {
	var b bytes.Buffer
	for i, ch := range s {
		if ch == '#' {
			b.Truncate(b.Len() - 2)
			d, _ := strconv.Atoi(s[i-2 : i])
			b.WriteString(string(96 + d))
		} else {
			b.WriteRune(ch + 48)
		}
	}
	return b.String()
}