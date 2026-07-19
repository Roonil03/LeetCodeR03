func rearrangeString(s string, x byte, y byte) string {
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}
	res := make([]byte, len(s))
	p := 0
	for i := range 28 {
		id := i
		if i == 26 {
			id = int(y - 'a')
		} else if i == 27 {
			id = int(x - 'a')
		} else if id == int(x-'a') || id == int(y-'a') {
			continue
		}
		ch := byte('a' + id)
		for freq[id] > 0 {
			res[p] = ch
			p++
			freq[id]--
		}
	}
	return string(res)
}