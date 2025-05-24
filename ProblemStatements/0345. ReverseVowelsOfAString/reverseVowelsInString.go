func reverseVowels(s string) string {
	vow := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	b := []byte(s)
	i, j := 0, len(b)-1
	for i < j {
		for i < j && !vow[b[i]] {
			i++
		}
		for i < j && !vow[b[j]] {
			j--
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b)
}