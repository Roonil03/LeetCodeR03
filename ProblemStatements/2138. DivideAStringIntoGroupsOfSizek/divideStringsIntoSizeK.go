func divideString(s string, k int, fill byte) []string {
	res := []string{}
	n := len(s)
	for i := 0; i < n; i += k {
		end := i + k
		if end <= n {
			res = append(res, s[i:end])
		} else {
			group := s[i:]
			for len(group) < k {
				group += string(fill)
			}
			res = append(res, group)
		}
	}
	return res
}