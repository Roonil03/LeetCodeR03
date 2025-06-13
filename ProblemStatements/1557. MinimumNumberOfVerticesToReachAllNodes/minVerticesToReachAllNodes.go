func findSmallestSetOfVertices(n int, edges [][]int) []int {
	id := make([]int, n)
	for _, e := range edges {
		id[e[1]]++
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if id[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}