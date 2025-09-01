func lengthLongestPath(input string) int {
	res, path := 0, make(map[int]int)
	path[0] = 0
	l := strings.Split(input, "\n")
	for _, i := range l {
		lvl := 0
		for strings.HasPrefix(i, "\t") {
			i = i[1:]
			lvl++
		}
		a := path[lvl] + len(i) + 1
		path[lvl+1] = a
		if strings.Contains(i, ".") {
			if a-1 > res {
				res = a - 1
			}
		}
	}
	return res
}