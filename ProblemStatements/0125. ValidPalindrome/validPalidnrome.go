func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isIt(rune(s[i])) {
			i++
		}
		for i < j && !isIt(rune(s[j])) {
			j--
		}
		if i < j {
			if toLower(rune(s[i])) != toLower(rune(s[j])) {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isIt(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}