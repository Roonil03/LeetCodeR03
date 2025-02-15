func compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}
	wi := 0
	i := 0
	for i < n {
		char := chars[i]
		count := 0
		for i < n && chars[i] == char {
			i++
			count++
		}
		chars[wi] = char
		wi++
		if count > 1 {
			cs := strconv.Itoa(count)
			for _, c := range cs {
				chars[wi] = byte(c)
				wi++
			}
		}
	}
	return wi
}