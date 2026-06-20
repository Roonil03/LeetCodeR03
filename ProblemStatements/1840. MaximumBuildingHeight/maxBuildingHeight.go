func maxBuilding(n int, restrictions [][]int) int {
	res := append(restrictions, []int{1, 0})
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	if res[len(res)-1][0] != n {
		res = append(res, []int{n, n - 1})
	}
	m := len(res)
	for i := 1; i < m; i++ {
		res[i][1] = min(res[i][1], res[i-1][1]+(res[i][0]-res[i-1][0]))
	}
	for i := m - 2; i >= 0; i-- {
		res[i][1] = min(res[i][1], res[i+1][1]+(res[i+1][0]-res[i][0]))
	}
	ans := 0
	for i := range m - 1 {
		id1, h1 := res[i][0], res[i][1]
		id2, h2 := res[i+1][0], res[i+1][1]
		peak := (h1 + h2 + (id2 - id1)) / 2
		ans = max(ans, peak)
	}
	return ans
}