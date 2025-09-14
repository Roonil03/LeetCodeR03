func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]struct{}, len(wordlist))
	lower := make(map[string]string, len(wordlist))
	mask := make(map[string]string, len(wordlist))
	devowel := func(s string) string {
		sb := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				sb[i] = '*'
			} else {
				sb[i] = c
			}
		}
		return string(sb)
	}
	for _, w := range wordlist {
		exact[w] = struct{}{}
		lw := strings.ToLower(w)
		if _, ok := lower[lw]; !ok {
			lower[lw] = w
		}
		mw := devowel(lw)
		if _, ok := mask[mw]; !ok {
			mask[mw] = w
		}
	}
	res := make([]string, len(queries))
	for i, q := range queries {
		if _, ok := exact[q]; ok {
			res[i] = q
			continue
		}
		lq := strings.ToLower(q)
		if w, ok := lower[lq]; ok {
			res[i] = w
			continue
		}
		mq := devowel(lq)
		if w, ok := mask[mq]; ok {
			res[i] = w
			continue
		}
		res[i] = ""
	}
	return res
}