func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	a, b := false, false
	for _, ch := range word {
		if isDigit(ch) {
			continue
		} else if isVowel(ch) {
			a = true
		} else if isOther(ch) {
			b = true
		} else {
			return false
		}
	}
	return a && b
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'O' || ch == 'I' || ch == 'U' || ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func isOther(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}