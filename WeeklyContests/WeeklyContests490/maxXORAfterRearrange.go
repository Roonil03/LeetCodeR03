func maximumXor(s string, t string) string {
	n := len(s)
	z := strings.Count(t, "0")
	o := n - z
	res := make([]byte, n)
	for i := range res {
		if s[i] == '0' {
			if o > 0 {
				res[i] = '1'
				o--
			} else {
				res[i] = '0'
				z--
			}
		} else {
			if z > 0 {
				res[i] = '1'
				z--
			} else {
				res[i] = '0'
				o--
			}
		}
	}
	return string(res)
}