func decodeString(s string) string {
	a := make([]int, 0)
	b := make([]string, 0)
	a1, b1 := 0, ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			a1 = a1*10 + int(ch-'0')
		} else if ch == '[' {
			a = append(a, a1)
			b = append(b, b1)
			a1 = 0
			b1 = ""
		} else if ch == ']' {
			n := a[len(a)-1]
			a = a[:len(a)-1]
			prev := b[len(b)-1]
			b = b[:len(b)-1]
			b1 = prev + strings.Repeat(b1, n)
		} else {
			b1 += string(ch)
		}
	}
	return b1
}