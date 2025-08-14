func strWithout3a3b(a int, b int) string {
	res := []byte{}
	for a > 0 || b > 0 {
		n := len(res)
		p := false
		if n >= 2 && res[n-2] == res[n-1] {
			p = (res[n-1] != 'a')
		} else {
			p = (a > b)
		}
		if p {
			res = append(res, 'a')
			a--
			if a > b && a > 0 {
				res = append(res, 'a')
				a--
			}
		} else {
			res = append(res, 'b')
			b--
			if b > a && b > 0 {
				res = append(res, 'b')
				b--
			}
		}
	}
	return string(res)
}