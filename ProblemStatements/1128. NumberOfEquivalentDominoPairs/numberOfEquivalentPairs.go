func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[[2]int]int)
	res := 0
	for _, i := range dominoes {
		a, b := i[0], i[1]
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		res += m[key]
		m[key]++
	}
	return res
}