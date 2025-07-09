func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	res, l := 0, math.MinInt32
	for _, p := range pairs {
		s, e := p[0], p[1]
		if l < s {
			res++
			l = e
		}
	}
	return res
}