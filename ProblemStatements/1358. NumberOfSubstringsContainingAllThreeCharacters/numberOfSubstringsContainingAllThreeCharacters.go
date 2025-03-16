func numberOfSubstrings(s string) int {
	n := len(s)
	var res int
	c := make([]int, 3)
	l := 0
	for r := 0; r < n; r++ {
		c[s[r]-'a']++
		for l <= r && c[0] > 0 && c[1] > 0 && c[2] > 0 {
			res += int(n - r)
			c[s[l]-'a']--
			l++
		}
	}
	return res
}