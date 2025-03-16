func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, h := 0, n
	for l < h {
		m := l + (h-l)/2
		if letters[m] > target {
			h = m
		} else {
			l = m + 1
		}
	}
	return letters[l%n]
}