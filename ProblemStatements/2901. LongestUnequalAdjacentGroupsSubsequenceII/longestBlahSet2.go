func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	codes := make([]uint64, n)
	for i, w := range words {
		var code uint64
		for j := 0; j < len(w); j++ {
			code |= uint64(w[j]-'a') << (5 * j)
		}
		codes[i] = code
	}
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxLen, maxIdx := 1, 0
	patternsByLen := make(map[int][]map[uint64][]int)
	for i, w := range words {
		L := len(w)
		if _, ok := patternsByLen[L]; !ok {
			buckets := make([]map[uint64][]int, L)
			for j := range buckets {
				buckets[j] = make(map[uint64][]int)
			}
			patternsByLen[L] = buckets
		}
		buckets := patternsByLen[L]
		gi := groups[i]
		codeI := codes[i]
		best, bp := 1, -1
		for j := 0; j < L; j++ {
			mask := ^(uint64(31) << (5 * j))
			pat := codeI & mask
			if lst, ok := buckets[j][pat]; ok {
				for _, k := range lst {
					if groups[k] != gi && dp[k]+1 > best {
						best = dp[k] + 1
						bp = k
					}
				}
			}
		}
		dp[i] = best
		prev[i] = bp
		if best > maxLen {
			maxLen = best
			maxIdx = i
		}
		for j := 0; j < L; j++ {
			mask := ^(uint64(31) << (5 * j))
			pat := codeI & mask
			buckets[j][pat] = append(buckets[j][pat], i)
		}
	}
	res := make([]string, 0, maxLen)
	for cur := maxIdx; cur != -1; cur = prev[cur] {
		res = append(res, words[cur])
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}