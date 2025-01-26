func maximumInvitations(favorite []int) int {
	n := len(favorite)
	d := make([]int, n)
	for _, f := range favorite {
		d[f]++
	}
	q := []int{}
	for i := range favorite {
		if d[i] == 0 {
			q = append(q, i)
		}
	}
	l := make([]int, n)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		f := favorite[i]
		l[f] = max(l[f], l[i]+1)
		d[f]--
		if d[f] == 0 {
			q = append(q, f)
		}
	}
	m := 0
	p := 0
	for i := range favorite {
		if d[i] == 0 {
			continue
		}
		c := 0
		k := i
		for d[k] > 0 {
			d[k] = 0
			k = favorite[k]
			c++
		}
		if c == 2 {
			p += 2 + l[i] + l[favorite[i]]
		} else {
			m = max(m, c)
		}
	}
	return max(m, p)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
