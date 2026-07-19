func transformStr(s string, strs []string) []bool {
	n := len(s)
	p := make([]int, n)
	z := 0
	for i := range n {
		if s[i] == '0' {
			z++
		}
		p[i] = z
	}
	r := make([]bool, len(strs))
	for i, v := range strs {
		vz, vo := 0, 0
		for j := range n {
			if v[j] == '0' {
				vz++
			} else if v[j] == '1' {
				vo++
			}
		}
		if vz > z || vo > n-z {
			continue
		}
		rem := z - vz
		c := 0
		fg := true
		for j := range n {
			if v[j] == '0' {
				c++
			} else if v[j] == '?' && rem > 0 {
				c++
				rem--
			}
			if c < p[j] {
				fg = false
				break
			}
		}
		r[i] = fg
	}
	return r
}