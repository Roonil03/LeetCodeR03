func createGrid(m int, n int) []string {
	res := make([]string, m)
	r := make([]byte, n)
	for i := range r {
		r[i] = '#'
	}
	r[0] = '.'
	s := string(r)
	for i := range m - 1 {
		res[i] = s
	}
	for i := range r {
		r[i] = '.'
	}
	res[m-1] = string(r)
	return res
}