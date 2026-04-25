func furthestDistanceFromOrigin(moves string) int {
	res := [3]int{}
	for _, r := range moves {
		switch r {
		case 'L':
			res[0]++
		case 'R':
			res[1]++
		case '_':
			res[2]++
		}
	}
	if res[0] > res[1] {
		return res[0] - res[1] + res[2]
	}
	return res[1] - res[0] + res[2]
}