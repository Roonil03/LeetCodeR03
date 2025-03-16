func findLadders(a string, b string, c []string) [][]string {
	d := make(map[string]bool)
	for _, e := range c {
		d[e] = true
	}
	if !d[b] {
		return [][]string{}
	}
	delete(d, a)
	f := make(map[string][]string)
	g := make(map[string]bool)
	g[a] = true
	h := false
	for len(g) > 0 && !h {
		for i := range g {
			delete(d, i)
		}
		j := make(map[string]bool)
		for i := range g {
			k := []byte(i)
			for l := 0; l < len(k); l++ {
				m := k[l]
				for n := byte('a'); n <= 'z'; n++ {
					k[l] = n
					o := string(k)
					if d[o] {
						j[o] = true
						f[o] = append(f[o], i)

						if o == b {
							h = true
						}
					}
				}
				k[l] = m
			}
		}
		g = j
	}
	if !h {
		return [][]string{}
	}
	return p(a, b, f)
}

func p(a string, b string, c map[string][]string) [][]string {
	if a == b {
		return [][]string{{a}}
	}
	var d [][]string
	for _, e := range c[b] {
		f := p(a, e, c)
		for _, g := range f {
			h := make([]string, len(g))
			copy(h, g)
			h = append(h, b)
			d = append(d, h)
		}
	}
	return d
}