func findMinMoves(machines []int) int {
	s := 0
	for _, v := range machines {
		s += v
	}
	if s%len(machines) != 0 {
		return -1
	}
	t := s / len(machines)
	res, c := 0, 0
	for _, v := range machines {
		v -= t
		c += v
		a := c
		if a < 0 {
			a = -a
		}
		res = max(a, res, v)
	}
	return res
}