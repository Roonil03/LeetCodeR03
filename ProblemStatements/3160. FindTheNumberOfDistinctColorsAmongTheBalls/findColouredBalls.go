func queryResults(limit int, queries [][]int) []int {
	col := make(map[int]int)
	freq := make(map[int]int)
	res := []int{}
	for _, q := range queries {
		x, y := q[0], q[1]
		if o1, exists := freq[x]; exists {
			col[o1]--
			if col[o1] == 0 {
				delete(col, o1)
			}
		}
		freq[x] = y
		col[y]++
		res = append(res, len(col))
	}
	return res
}