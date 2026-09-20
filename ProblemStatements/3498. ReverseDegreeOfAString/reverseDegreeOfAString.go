func reverseDegree(s string) int {
	n := len(s)
	res := 123 * n * (n + 1) / 2
	for i, v := range s {
		res -= int(v) * (i + 1)
	}
	return res
}