func areAlmostEqual(s1 string, s2 string) bool {
	a, b := make([]int, 26), make([]int, 26)
	cnt := 0
	for i := 0; i < len(s1); i++ {
		a[s1[i]-'a']++
		b[s2[i]-'a']++
		if s1[i] != s2[i] {
			cnt++
		}
	}
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return cnt <= 2
}

