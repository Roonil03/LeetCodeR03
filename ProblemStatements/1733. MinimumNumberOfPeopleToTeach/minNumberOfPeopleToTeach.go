func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	n1 := len(languages)
	sets := make([]map[int]struct{}, n1+1)
	for i := 1; i <= n1; i++ {
		s := make(map[int]struct{}, len(languages[i-1]))
		for _, l := range languages[i-1] {
			s[l] = struct{}{}
		}
		sets[i] = s
	}
	need := make(map[int]struct{})
	for _, f := range friendships {
		u, v := f[0], f[1]
		ok := false
		if len(sets[u]) < len(sets[v]) {
			for l := range sets[u] {
				if _, t := sets[v][l]; t {
					ok = true
					break
				}
			}
		} else {
			for l := range sets[v] {
				if _, t := sets[u][l]; t {
					ok = true
					break
				}
			}
		}
		if !ok {
			need[u] = struct{}{}
			need[v] = struct{}{}
		}
	}
	if len(need) == 0 {
		return 0
	}
	count := make([]int, n+1)
	for u := range need {
		for l := range sets[u] {
			count[l]++
		}
	}
	m := 0
	for l := 1; l <= n; l++ {
		if count[l] > m {
			m = count[l]
		}
	}
	return len(need) - m
}