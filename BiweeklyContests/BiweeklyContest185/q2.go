func minLights(lights []int) int {
	n := len(lights)
	d := make([]int, n+1)
	for i, v := range lights {
		if v > 0 {
			l, r := max(i-v, 0), min(n, i+v+1)
			d[l]++
			d[r]--
		}
	}
	a, b, c := 0, 0, 0
	for i := range n {
		c += d[i]
		if c == 0 {
			b++
		} else if b > 0 {
			a += (b + 2) / 3
			b = 0
		}
	}
	return a + (b+2)/3
}