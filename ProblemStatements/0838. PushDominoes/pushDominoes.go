func pushDominoes(dominoes string) string {
	type sym struct {
		id   int
		char byte
	}
	syms := []sym{}
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i] != '.' {
			syms = append(syms, sym{i, dominoes[i]})
		}
	}
	syms = append([]sym{{-1, 'L'}}, syms...)
	syms = append(syms, sym{len(dominoes), 'R'})
	res := []byte(dominoes)
	for i := 0; i < len(syms)-1; i++ {
		l, r := syms[i], syms[i+1]
		if l.char == r.char {
			for k := l.id + 1; k < r.id; k++ {
				res[k] = l.char
			}
		} else if l.char == 'R' && r.char == 'L' {
			l1, r1 := l.id+1, r.id-1
			for l1 < r1 {
				res[l1] = 'R'
				res[r1] = 'L'
				l1++
				r1--
			}
		}
	}
	return string(res)
}