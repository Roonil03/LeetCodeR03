func validIPAddress(queryIP string) string {
	d, c := 0, 0
	n := len(queryIP)
	for i := range n {
		if queryIP[i] == '.' {
			d++
		} else if queryIP[i] == ':' {
			c++
		}
	}
	if d == 3 {
		for i, j := 0, 0; i <= n; i++ {
			if i == n || queryIP[i] == '.' {
				if i-j < 1 || i-j > 3 || (i-j > 1 && queryIP[j] == '0') {
					return "Neither"
				}
				val := 0
				for k := j; k < i; k++ {
					if queryIP[k] < '0' || queryIP[k] > '9' {
						return "Neither"
					}
					val = val*10 + int(queryIP[k]-'0')
				}
				if val > 255 {
					return "Neither"
				}
				j = i + 1
			}
		}
		return "IPv4"
	}
	if c == 7 {
		for i, j := 0, 0; i <= n; i++ {
			if i == n || queryIP[i] == ':' {
				if i-j < 1 || i-j > 4 {
					return "Neither"
				}
				for k := j; k < i; k++ {
					ch := queryIP[k]
					if (ch < '0' || ch > '9') && ((ch|32) < 'a' || (ch|32) > 'f') {
						return "Neither"
					}
				}
				j = i + 1
			}
		}
		return "IPv6"
	}
	return "Neither"
}