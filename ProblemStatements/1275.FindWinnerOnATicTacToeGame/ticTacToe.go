func tictactoe(moves [][]int) string {
	rows := [3]int{}
	cols := [3]int{}
	d := 0
	ad := 0
	// w := 1
	for i, m := range moves {
		r, c := m[0], m[1]
		add := 1
		if i%2 == 1 {
			add = -1
		}
		rows[r] += add
		cols[c] += add
		if r == c {
			d += add
		}
		if r+c == 2 {
			ad += add
		}
		if abs(rows[r]) == 3 || abs(cols[c]) == 3 || abs(d) == 3 || abs(ad) == 3 {
			if add == 1 {
				return "A"
			} else {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}