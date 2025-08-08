func soupServings(n int) float64 {
	if n > 4800 {
		return 1
	}
	m := (n + 24) / 25
	mem := map[[2]int]float64{}
	var dfs func(int, int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1.0
		}
		if b <= 0 {
			return 0.0
		}
		key := [2]int{a, b}
		if v, ok := mem[key]; ok {
			return v
		}
		v := 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		mem[key] = v
		return v
	}
	return dfs(m, m)
}