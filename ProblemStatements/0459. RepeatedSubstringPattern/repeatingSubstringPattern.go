func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	for i := n / 2; i >= 1; i-- {
		if n%i == 0 {
			if help(s, i) {
				return true
			}
		}
	}
	return false
}

func help(s string, subLen int) bool {
	p := s[:subLen]
	for i := subLen; i < len(s); i += subLen {
		if s[i:i+subLen] != p {
			return false
		}
	}
	return true
}