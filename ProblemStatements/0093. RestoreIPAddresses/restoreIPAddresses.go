func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	n, e := strconv.Atoi(s)
	return e == nil && n >= 0 && n <= 255
}

func backtrack(s string, i int, p []string, r *[]string) {
	if len(p) == 4 {
		if i == len(s) {
			*r = append(*r, strings.Join(p, "."))
		}
		return
	}
	for j := 1; j <= 3; j++ {
		if i+j > len(s) {
			break
		}
		t := s[i : i+j]
		if isValid(t) {
			backtrack(s, i+j, append(p, t), r)
		}
	}
}

func restoreIpAddresses(s string) []string {
	var r []string
	backtrack(s, 0, []string{}, &r)
	return r
}