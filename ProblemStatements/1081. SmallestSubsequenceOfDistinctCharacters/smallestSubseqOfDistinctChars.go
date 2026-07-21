func smallestSubsequence(s string) string {
	a := make([]int, 256)
	for i := range s {
		a[s[i]] = i
	}
	st := []byte{}
	vis := make([]bool, 236)
	for i := range s {
		ch := s[i]
		if vis[ch] {
			continue
		}
		for len(st) > 0 && st[len(st)-1] > ch && a[st[len(st)-1]] > i {
			vis[st[len(st)-1]] = false
			st = st[:len(st)-1]
		}
		st = append(st, ch)
		vis[ch] = true
	}
	return string(st)
}