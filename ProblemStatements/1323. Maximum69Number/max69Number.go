func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '6' {
			b[i] = '9'
			res, _ := strconv.Atoi(string(b))
			return res
		}
	}
	// b[len(b) - 1] = '9'
	// res, _ := strconv.Atoi(string(b))
	return num
}