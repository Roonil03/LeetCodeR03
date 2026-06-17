func processStr(s string, k int64) byte {
	var m int64
	for i := range s {
		if s[i] == '*' {
			if m > 0 {
				m--
			}
		} else if s[i] == '#' {
			m <<= 1
		} else if s[i] != '%' {
			m++
		}
	}
	if k >= m {
		return '.'
	}
	for i := len(s) - 1; ; i-- {
		if s[i] == '*' {
			m++
		} else if s[i] == '#' {
			m /= 2
			if k >= m {
				k -= m
			}
		} else if s[i] == '%' {
			k = m - 1 - k
		} else {
			m--
			if k == m {
				return s[i]
			}
		}
	}
}