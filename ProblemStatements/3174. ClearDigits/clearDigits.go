func clearDigits(st string) string {
	s := []rune{}
	for _, ch := range st {
		if unicode.IsDigit(ch) {
			s = s[:len(s)-1]
		} else {
			s = append(s, ch)
		}
	}
	return string(s)
}