func scoreOfParentheses(s string) int {
	res, t := 0, 0
	for i := range s {
		if s[i] == '(' {
			t++
		} else {
			t--
			if s[i-1] == '(' {
				res += 1 << t
			}
		}
	}
	return res
}