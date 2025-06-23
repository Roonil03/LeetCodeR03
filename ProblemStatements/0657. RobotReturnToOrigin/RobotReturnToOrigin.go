func judgeCircle(moves string) bool {
	a, b := 0, 0
	for _, c := range moves {
		switch c {
		case 'U':
			a++
		case 'D':
			a--
		case 'R':
			b++
		case 'L':
			b--
		}
	}
	return a == 0 && b == 0
}