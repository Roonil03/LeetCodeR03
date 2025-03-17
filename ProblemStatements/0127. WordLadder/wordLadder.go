func ladderLength(beginWord string, endWord string, wordList []string) int {
	wm := make(map[string]bool)
	for _, n := range wordList {
		wm[n] = true
	}
	if !wm[endWord] {
		return 0
	}
	q := []string{beginWord}
	lvl := 1
	vis := make(map[string]bool)
	vis[beginWord] = true
	for len(q) > 0 {
		cl := len(q)
		for i := 0; i < cl; i++ {
			a := q[0]
			q = q[1:]
			if a == endWord {
				return lvl
			}
			b := []rune(a)
			for j := 0; j < len(b); j++ {
				o := b[j]
				for c := 'a'; c <= 'z'; c++ {
					if c == o {
						continue
					}
					b[j] = c
					n := string(b)
					if wm[n] && !vis[n] {
						q = append(q, n)
						vis[n] = true
					}
				}
				b[j] = o
			}
		}
		lvl++
	}
	return 0
}