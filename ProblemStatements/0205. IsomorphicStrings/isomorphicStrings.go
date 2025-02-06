func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := range m1 {
		m1[i] = -1
		m2[i] = -1
	}
	for i := 0; i < len(s); i++ {
		c1, c2 := int(s[i]), int(t[i])
		if m1[c1] == -1 && m2[c2] == -1 {
			m1[c1] = c2
			m2[c2] = c1
		} else if m1[c1] != c2 || m2[c2] != c1 {
			return false
		}
	}
	return true
}