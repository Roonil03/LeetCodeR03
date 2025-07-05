func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 32; i++ {
		abit := (a >> i) & 1
		bbit := (b >> i) & 1
		cbit := (c >> i) & 1
		if cbit == 1 {
			if abit|bbit == 0 {
				res++
			}
		} else {
			if abit == 1 {
				res++
			}
			if bbit == 1 {
				res++
			}
		}
	}
	return res
}