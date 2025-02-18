func smallestNumber(pattern string) string {
	n := len(pattern)
	res := make([]byte, n+1)
	s := make([]byte, 0)
	c := byte('1')
	for i := 0; i <= n; i++ {
		s = append(s, c)
		c++
		if i == n || pattern[i] == 'I' {
			for len(s) > 0 {
				top := len(s) - 1
				res[i-top] = s[top]
				s = s[:top]
			}
		}
	}
	return string(res)
}