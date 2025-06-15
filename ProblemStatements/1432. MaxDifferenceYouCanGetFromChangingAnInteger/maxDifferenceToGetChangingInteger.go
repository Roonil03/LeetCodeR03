func maxDiff(num int) int {
	s := strconv.Itoa(num)
	m1, m2 := num, num
	for i := 0; i < len(s); i++ {
		x := s[i]
		for y := '0'; y <= '9'; y++ {
			nw := ""
			val := true
			for j := 0; j < len(s); j++ {
				if s[j] == x {
					if j == 0 && y == '0' {
						val = false
						break
					}
					nw += string(y)
				} else {
					nw += string(s[j])
				}
			}
			if val {
				a, _ := strconv.Atoi(nw)
				if a > m2 {
					m2 = a
				}
				if a < m1 && a != 0 {
					m1 = a
				}
			}
		}
	}
	return m2 - m1
}