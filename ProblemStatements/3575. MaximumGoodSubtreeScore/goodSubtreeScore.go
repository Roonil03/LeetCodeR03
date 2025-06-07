func goodSubtreeSum(vals []int, par []int) int {
	const MOD = 1e9 + 7
	n := len(vals)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[par[i]] = append(adj[par[i]], i)
	}
	ans := 0
	var dfs func(u int) map[int]int
	dfs = func(u int) map[int]int {
		du := map[int]int{0: 0}
		v := vals[u]
		if v < 0 {
			v = -v
		}
		s := strconv.Itoa(v)
		d := make(map[byte]bool)
		ok := true
		for i := 0; i < len(s); i++ {
			if d[s[i]] {
				ok = false
				break
			}
			d[s[i]] = true
		}
		if ok {
			m := 0
			for i := 0; i < len(s); i++ {
				m |= 1 << (s[i] - '0')
			}
			du[m] = vals[u]
		}
		for _, ch := range adj[u] {
			dv := dfs(ch)
			cp := make(map[int]int)
			for k, v := range du {
				cp[k] = v
			}
			for mv, sv := range dv {
				for mu, su := range cp {
					if mu&mv == 0 {
						nm := mu | mv
						if cur, ex := du[nm]; !ex || su+sv > cur {
							du[nm] = su + sv
						}
					}
				}
			}
		}
		mx := 0
		for _, val := range du {
			if val > mx {
				mx = val
			}
		}
		ans += mx
		return du
	}
	dfs(0)
	return ans % MOD
}