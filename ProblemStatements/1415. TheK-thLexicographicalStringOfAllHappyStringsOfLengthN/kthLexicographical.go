func getHappyString(n int, k int) string {
	t := 3 << (n - 1 - 0)
	if k > t {
		return ""
	}
	var res string
	var bt func(p string, id *int) bool
	bt = func(p string, id *int) bool {
		if len(p) == n {
			*id--
			if *id == 0 {
				res = p
				return true
			}
			return false
		}
		c := []rune{'a', 'b', 'c'}
		if len(p) > 0 {
			temp := []rune{}
			last := rune(p[len(p)-1])
			for _, ch := range c {
				if ch != last {
					temp = append(temp, ch)
				}
			}
			c = temp
		}
		rem := n - len(p) - 1
		b := 1
		if rem > 0 {
			b = 1 << rem
		}
		for _, ch := range c {
			if *id > b {
				*id -= b
			} else {
				if bt(p+string(ch), id) {
					return true
				}
			}
		}
		return false
	}
	bt("", &k)
	return res
}