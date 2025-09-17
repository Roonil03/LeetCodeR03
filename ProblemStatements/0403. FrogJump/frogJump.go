func canCross(stones []int) bool {
	if len(stones) > 1 && stones[1] != 1 {
		return false
	}
	sp := make(map[int]int, len(stones))
	for i, pos := range stones {
		sp[pos] = i
	}
	type a struct {
		id   int
		jump int
	}
	mem := make(map[a]bool)
	var dfs func(int, int) bool
	dfs = func(cur, ljump int) bool {
		if cur == len(stones)-1 {
			return true
		}
		s := a{id: cur, jump: ljump}
		if val, found := mem[s]; found {
			return val
		}
		for nj := ljump - 1; nj <= ljump+1; nj++ {
			if nj > 0 {
				np := stones[cur] + nj
				if ni, found := sp[np]; found {
					if dfs(ni, nj) {
						mem[s] = true
						return true
					}
				}
			}
		}
		mem[s] = false
		return false
	}
	return dfs(0, 0)
}