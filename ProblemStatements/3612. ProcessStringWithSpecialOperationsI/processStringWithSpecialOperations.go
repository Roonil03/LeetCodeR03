func processStr(s string) string {
	n := len(s)
	cap := 2 * n
	if cap < 32 {
		cap = 32
	}
	buf := make([]byte, cap)
	st := cap / 2
	end := st
	fg := false
	for i := range n {
		ch := s[i]
		switch ch {
		case '*':
			if end > st {
				if !fg {
					end--
				} else {
					st++
				}
			}
		case '#':
			l := end - st
			if l > 0 {
				if end+l > len(buf) {
					buf, st, end = h1(buf, st, end, 0, l)
				}
				copy(buf[end:], buf[st:end])
				end += l
			}
		case '%':
			fg = !fg
		default:
			if !fg {
				if end >= len(buf) {
					buf, st, end = h1(buf, st, end, 0, 1)
				}
				buf[end] = ch
				end++
			} else {
				if st <= 0 {
					buf, st, end = h1(buf, st, end, 1, 0)
				}
				st--
				buf[st] = ch
			}
		}
	}
	r := buf[st:end]
	if fg {
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	}
	return string(r)
}

func h1(buf []byte, st, end, l, r int) ([]byte, int, int) {
	cur := end - st
	cap := len(buf) * 2
	if cap < cur+l+r {
		cap = cur + l + r + 16
	}
	b := make([]byte, cap)
	s := (cap - cur) / 2
	copy(b[s:], buf[st:end])
	return b, s, s + cur
}