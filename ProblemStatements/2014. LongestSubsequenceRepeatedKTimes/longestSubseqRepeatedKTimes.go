func longestSubsequenceRepeatedK(s string, k int) string {
	n0 := len(s)
	arr := make([]int, n0)
	for i := range s {
		arr[i] = int(s[i] - 'a')
	}
	var cnt [26]int
	for _, c := range arr {
		cnt[c]++
	}
	cand := []int{}
	for c := 0; c < 26; c++ {
		if cnt[c] >= k {
			cand = append(cand, c)
		}
	}
	if len(cand) == 0 {
		return ""
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cand)))
	filtered := []int{}
	for _, c := range arr {
		if cnt[c] >= k {
			filtered = append(filtered, c)
		}
	}
	n := len(filtered)
	nxt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		nxt[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			nxt[i][j] = n
		}
	}
	for i := n - 1; i >= 0; i-- {
		copy(nxt[i], nxt[i+1])
		nxt[i][filtered[i]] = i
	}
	maxL := n / k
	res := []int{}
	var best []int
	var dfs func(d, L int) bool
	dfs = func(d, L int) bool {
		if d == L {
			best = append([]int(nil), res...)
			return true
		}
		for _, c := range cand {
			res = append(res, c)
			pos := 0
			ok := true
			for t := 0; t < k && ok; t++ {
				for _, x := range res {
					pos = nxt[pos][x] + 1
					if pos > n {
						ok = false
						break
					}
				}
			}
			if ok && dfs(d+1, L) {
				return true
			}
			res = res[:len(res)-1]
		}
		return false
	}
	for L := maxL; L > 0; L-- {
		res = res[:0]
		if dfs(0, L) {
			break
		}
	}
	var sb strings.Builder
	for _, c := range best {
		sb.WriteByte(byte('a' + c))
	}
	return sb.String()
}